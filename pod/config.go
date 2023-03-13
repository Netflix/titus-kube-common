package pod

import (
	"errors"
	"regexp"
	"time"

	resourceCommon "github.com/Netflix/titus-kube-common/resource"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

// Config contains configuration parameters parsed out from various places in the pod
// (such as annotations). All fields are pointers, to differentiate between a field being
// unset and the empty value.
type Config struct {
	AssignIPv6Address        *bool
	AccountID                *string
	AppArmorProfile          *string
	CapacityGroup            *string
	CPUBurstingEnabled       *bool
	ContainerInfo            *string
	EgressBandwidth          *resource.Quantity
	ElasticIPPool            *string
	ElasticIPs               *string
	EntrypointShellSplitting *bool
	FuseEnabled              *bool
	HostnameStyle            *string
	IAMRole                  *string
	InjectedEnvVarNames      []string
	IngressBandwidth         *resource.Quantity
	IMDSRequireToken         *string
	JobAcceptedTimestampMs   *uint64
	JobDescriptor            *string
	JobID                    *string
	JobType                  *string
	JumboFramesEnabled       *bool
	KvmEnabled               *bool
	LogKeepLocalFile         *bool
	LogUploadCheckInterval   *time.Duration
	LogUploadThresholdTime   *time.Duration
	LogUploadRegExp          *regexp.Regexp
	LogStdioCheckInterval    *time.Duration
	LogS3WriterIAMRole       *string
	LogS3BucketName          *string
	LogS3PathPrefix          *string
	NetworkMode              *string
	NetworkBurstingEnabled   *bool
	NflxIMDSEnabled          *bool
	OomScoreAdj              *int32
	PodSchemaVersion         *uint32
	ResourceCPU              *resource.Quantity
	ResourceDisk             *resource.Quantity
	ResourceGPU              *resource.Quantity
	ResourceMemory           *resource.Quantity
	ResourceNetwork          *resource.Quantity
	SchedPolicy              *string
	SeccompAgentNetEnabled   *bool
	SeccompAgentPerfEnabled  *bool
	TrafficSteeringEnabled   *bool
	SecurityGroupIDs         *[]string
	StaticIPAllocationUUID   *string
	SystemEnvVarNames        []string
	SubnetIDs                *[]string
	TaskID                   *string
	TTYEnabled               *bool
	WorkloadDetail           *string
	WorkloadName             *string
	WorkloadMetadata         *string
	WorkloadMetadataSig      *string
	WorkloadOwnerEmail       *string
	WorkloadSequence         *string
	WorkloadStack            *string
}

// Sidecar represents a sidecar that's configured to run as part of the container
type Sidecar struct {
	Enabled bool
	Image   string
	Name    string
	Version int
}

// PodToConfig pulls out values from a pod and turns them into a Config
func PodToConfig(pod *corev1.Pod) (*Config, error) {
	pConf := &Config{}

	err := parseAnnotations(pod, pConf)
	if err != nil {
		return pConf, err
	}

	err = parseLabels(pod, pConf)
	if err != nil {
		return pConf, err
	}

	err = parsePodFields(pod, pConf)
	if err != nil {
		return pConf, err
	}

	return pConf, err
}

func parsePodFields(pod *corev1.Pod, pConf *Config) error {
	mainContainer := GetMainUserContainer(pod)
	if mainContainer == nil {
		return errors.New("could not find main container in pod")
	}

	resources := mainContainer.Resources.Limits
	pConf.ResourceCPU = resourcePtr(resources, corev1.ResourceCPU)
	pConf.ResourceDisk = resourcePtr(resources, corev1.ResourceEphemeralStorage)
	pConf.ResourceGPU = resourcePtr(resources, resourceCommon.ResourceNameGpu)
	pConf.ResourceMemory = resourcePtr(resources, corev1.ResourceMemory)
	pConf.ResourceNetwork = resourcePtr(resources, resourceCommon.ResourceNameNetwork)
	// XXX: do we need the legacy gpu and network resource names, too?

	if mainContainer.TTY {
		ttyEnabled := true
		pConf.TTYEnabled = &ttyEnabled
	}

	return nil
}

func resourcePtr(resources corev1.ResourceList, resName corev1.ResourceName) *resource.Quantity {
	res, ok := resources[resName]
	if !ok {
		return nil
	}

	return &res
}
