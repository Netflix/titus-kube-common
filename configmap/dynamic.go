package configmap

import (
	"context"
	"sync"
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	metaV1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/informers"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/kubernetes/typed/core/v1"
	"k8s.io/client-go/rest"
	"k8s.io/client-go/tools/cache"
)

type ConfigState struct {
	Raw       map[string]string
	Mapped    interface{}
	Timestamp time.Time
}

type ConfigMapperInput struct {
	RawCurrent map[string]string
	Previous   ConfigState
}

// ConfigMapper parses configmap data and returns user provided object or error
type ConfigMapper func(rawCurrent map[string]string, previous *ConfigState) (interface{}, error)

type OnUpdateCallback func(current *ConfigState, previous *ConfigState)

type OnErrorCallback func(err error, previous *ConfigState)

type Options struct {
	// If set to true, load data during initialization only, and do not refresh them.
	DisableDynamicUpdates bool
	// Update callback is called on each data evaluation both successful and failed
	OnUpdateCallback OnUpdateCallback
	// Called when there is an during config data processing
	OnErrorCallback OnErrorCallback
	// Configuration template. If set, and configmap is not found, a new one is created using this template.
	// If configmap is there, but some parameters are missing, there are added with the default values.
	Template map[string]string
	// If set to true, unrecognized parameters are removed from a configmap
	RemoveUnknownProperties bool
}

type DynamicConfig interface {
	// GetBootstrap returns the configuration state loaded on bootstrap. As bootstrap state never changes.
	GetBootstrap() interface{}

	// GetRawBootstrap returns the configmap loaded on bootstrap. As bootstrap state never changes.
	GetRawBootstrap() map[string]string

	// Get returns the latest configuration state
	Get() interface{}

	// GetRaw returns the latest configmap
	GetRaw() map[string]string

	// Close stops the internal refresh process
	Close()
}

type dynamicConfigInternal struct {
	// Parameters
	configMapName string
	options       Options
	configMapper  ConfigMapper
	// Internal
	client        *kubernetes.Clientset
	lock          sync.Mutex
	refreshCancel chan struct{}
	// Recently processed data
	bootstrap *ConfigState
	current   *ConfigState
}

func NewDynamicConfig(config *rest.Config, configMapName string, configMapper ConfigMapper, options Options) (DynamicConfig, error) {
	internal := dynamicConfigInternal{
		client:        kubernetes.NewForConfigOrDie(config),
		configMapName: configMapName,
		options:       options,
		configMapper:  configMapper,
	}
	data, err := internal.loadAndTryFixConfig()
	if err != nil {
		return nil, err
	}

	if err = internal.process(data); err != nil {
		return nil, err
	}
	internal.bootstrap = internal.current

	if !options.DisableDynamicUpdates {
		internal.startRefreshProcess(kubernetes.NewForConfigOrDie(config))
	}

	return &internal, nil
}

func NewDynamicConfigFromMap(data map[string]string, configMapper ConfigMapper, options Options) (DynamicConfig, error) {
	internal := dynamicConfigInternal{
		options:      options,
		configMapper: configMapper,
	}
	if err := internal.process(data); err != nil {
		return nil, err
	}
	internal.bootstrap = internal.current
	return &internal, nil
}

func (d *dynamicConfigInternal) GetBootstrap() interface{} {
	return d.bootstrap.Mapped
}

func (d *dynamicConfigInternal) GetRawBootstrap() map[string]string {
	return d.bootstrap.Raw
}

func (d *dynamicConfigInternal) Get() interface{} {
	return d.getCurrent().Mapped
}

func (d *dynamicConfigInternal) GetRaw() map[string]string {
	return d.getCurrent().Raw
}

func (d *dynamicConfigInternal) Close() {
	if d.refreshCancel != nil {
		d.refreshCancel <- struct{}{}
	}
}

func (d *dynamicConfigInternal) getCurrent() *ConfigState {
	d.lock.Lock()
	defer d.lock.Unlock()
	return d.current
}

func (d *dynamicConfigInternal) loadAndTryFixConfig() (map[string]string, error) {
	configMapsClient := d.client.CoreV1().ConfigMaps("default")
	controllerConfig, err := configMapsClient.Get(context.TODO(), d.configMapName, metaV1.GetOptions{})
	if err != nil {
		if d.options.Template == nil {
			return nil, err
		}
		if serr, ok := err.(*errors.StatusError); !ok || serr.ErrStatus.Reason != metaV1.StatusReasonNotFound {
			return nil, err
		}
		configMap := corev1.ConfigMap{
			ObjectMeta: metaV1.ObjectMeta{
				Namespace: "default",
				Name:      d.configMapName,
			},
			Data: d.options.Template,
		}
		createdConfigMap, createErr := configMapsClient.Create(context.TODO(), &configMap, metaV1.CreateOptions{})
		if createErr != nil {
			return nil, createErr
		}
		return createdConfigMap.Data, err
	}
	data := d.updateConfigMap(configMapsClient, controllerConfig)
	return data, nil
}

func (d *dynamicConfigInternal) updateConfigMap(configMapsClient v1.ConfigMapInterface,
	stored *corev1.ConfigMap) map[string]string {
	if d.options.Template == nil {
		return stored.Data
	}
	merged := map[string]string{}
	for key, value := range stored.Data {
		if d.options.RemoveUnknownProperties {
			if _, ok := d.options.Template[key]; ok {
				merged[key] = value
			}
		} else {
			merged[key] = value
		}
	}
	for key, value := range d.options.Template {
		if _, ok := stored.Data[key]; !ok {
			merged[key] = value
		}
	}

	stored.Data = merged
	updated, err := configMapsClient.Update(context.TODO(), stored, metaV1.UpdateOptions{})
	if err != nil {
		if d.options.OnErrorCallback != nil {
			d.options.OnErrorCallback(err, nil)
		}
		return stored.Data
	}

	return updated.Data
}

func (d *dynamicConfigInternal) startRefreshProcess(clientSet *kubernetes.Clientset) {
	informerFactory := informers.NewSharedInformerFactoryWithOptions(clientSet, time.Minute)
	configMapInformer := informerFactory.Core().V1().ConfigMaps().Informer()
	configMapInformer.AddEventHandler(cache.ResourceEventHandlerFuncs{
		AddFunc: func(obj interface{}) {
			d.processInformerUpdate(obj)
		},
		UpdateFunc: func(oldObj interface{}, obj interface{}) {
			d.processInformerUpdate(obj)
		},
	})

	informerFactory.Start(d.refreshCancel)
}

func (d *dynamicConfigInternal) processInformerUpdate(obj interface{}) {
	config := obj.(*corev1.ConfigMap)
	if config.Name == d.configMapName {
		_ = d.process(config.Data)
	}
}

func (d *dynamicConfigInternal) process(raw map[string]string) error {
	mapped, err := d.configMapper(raw, d.getCurrent())

	if err != nil {
		if d.options.OnUpdateCallback != nil {
			d.options.OnErrorCallback(err, d.current)
		}
		return err
	}
	current := ConfigState{
		Raw:       raw,
		Mapped:    mapped,
		Timestamp: time.Now(),
	}
	if d.options.OnUpdateCallback != nil {
		d.options.OnUpdateCallback(&current, d.current)
	}

	d.lock.Lock()
	defer d.lock.Unlock()
	d.current = &current

	return nil
}
