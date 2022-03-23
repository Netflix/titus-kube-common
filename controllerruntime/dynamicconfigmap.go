package controllerruntime

import (
	"context"
	"strconv"
	"sync"

	"github.com/go-logr/logr"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/predicate"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

var (
	_ reconcile.Reconciler = (*DynamicConfigMap)(nil)
)

type DynamicConfigMap struct {
	configMapName string
	client        client.Client
	log           logr.Logger
	rawConfig     map[string]string
	configMutex   sync.RWMutex
}

func NewDynamicConfigMap(ctx context.Context, mgr manager.Manager, configMapName string) (*DynamicConfigMap, error) {
	log := mgr.GetLogger().WithName("dynamic-config-map-" + configMapName)

	config, err := preloadConfig(ctx, log, mgr.GetAPIReader(), configMapName)
	if err != nil {
		return nil, err
	}

	d := &DynamicConfigMap{
		client:        mgr.GetClient(),
		log:           log,
		configMapName: configMapName,
	}
	d.setRawConfig(config)

	if err := d.setupWithManager(mgr); err != nil {
		return nil, err
	}

	return d, nil
}

func (c *DynamicConfigMap) Reconcile(ctx context.Context, req reconcile.Request) (reconcile.Result, error) {
	cm := &corev1.ConfigMap{}
	err := c.client.Get(ctx, req.NamespacedName, cm)

	if err == nil {
		c.log.Info("setting config", "data", cm.Data)
		c.setRawConfig(cm.Data)
		return reconcile.Result{}, nil
	}

	if apierrors.IsNotFound(err) {
		c.log.Info("config map not found, empty config set")
		c.setRawConfig(map[string]string{})
		return reconcile.Result{}, nil
	}

	return reconcile.Result{}, err
}

func (c *DynamicConfigMap) setupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&corev1.ConfigMap{}).
		WithEventFilter(predicate.Funcs{
			CreateFunc: func(event event.CreateEvent) bool {
				return event.Object.GetName() == c.configMapName
			},
			DeleteFunc: func(event event.DeleteEvent) bool {
				return event.Object.GetName() == c.configMapName
			},
			UpdateFunc: func(event event.UpdateEvent) bool {
				return event.ObjectNew.GetName() == c.configMapName
			},
			GenericFunc: func(event event.GenericEvent) bool {
				return event.Object.GetName() == c.configMapName
			},
		}).
		Complete(c)
}

func preloadConfig(ctx context.Context, log logr.Logger, client client.Reader, configMapName string) (map[string]string, error) {
	log = log.WithValues("name", configMapName)

	cm := &corev1.ConfigMap{}
	err := client.Get(ctx, types.NamespacedName{Namespace: "default", Name: configMapName}, cm)
	if err == nil {
		return cm.Data, nil
	}

	if apierrors.IsNotFound(err) {
		log.Info(configMapName + "config map not found, empty config initialized")
		return map[string]string{}, nil
	}
	log.Error(err, "failed to preload config map "+configMapName)
	return map[string]string{}, err
}

func (c *DynamicConfigMap) setRawConfig(m map[string]string) {
	c.configMutex.Lock()
	defer c.configMutex.Unlock()

	c.rawConfig = m
}

func (c *DynamicConfigMap) getRawValue(key string) (string, bool) {
	c.configMutex.RLock()
	defer c.configMutex.RUnlock()
	value, ok := c.rawConfig[key]
	return value, ok
}

func (c *DynamicConfigMap) GetBoolean(key string, dflt bool) bool {
	bRaw, ok := c.getRawValue(key)
	if !ok {
		return dflt
	}
	// Ignore errors.  Parse failure returns false
	b, _ := strconv.ParseBool(bRaw)
	return b
}

func (c *DynamicConfigMap) GetString(key string, dflt string) string {
	bRaw, ok := c.getRawValue(key)
	if !ok {
		return dflt
	}
	return bRaw
}

func (c *DynamicConfigMap) GetFloat64(key string, dflt float64) float64 {
	if valueStr, ok := c.getRawValue(key); ok {
		if value, err := strconv.ParseFloat(valueStr, 64); err == nil {
			return value
		}
	}
	return dflt
}

func (c *DynamicConfigMap) GetInt(key string, dflt int) int {
	if valueStr, ok := c.getRawValue(key); ok {
		if value, err := strconv.Atoi(valueStr); err == nil {
			return value
		}
	}
	return dflt
}

func (c *DynamicConfigMap) GetInt64(key string, dflt int64) int64 {
	if valueStr, ok := c.getRawValue(key); ok {
		if value, err := strconv.ParseInt(valueStr, 10, 64); err == nil {
			return value
		}
	}
	return dflt
}
