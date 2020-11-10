package pod

import (
	"time"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
)

// Config contains configuration parameters parsed out from various places in the pod
// (such as annotations). All fields are pointers, to differentiate between a field being
// unset and the empty value.
type Config struct {
	AssignIPv6Address      *bool
	AccountID              *string
	AppDetail              *string
	AppName                *string
	AppMetadata            *string
	AppMetadataSig         *string
	AppOwnerEmail          *string
	AppSequence            *string
	AppStack               *string
	BytesEnabled           *bool
	CapacityGroup          *string
	CPUBurstingEnabled     *bool
	ContainerInfo          *string
	EgressBandwidth        *resource.Quantity
	ElasticIPPool          *string
	ElasticIPs             *string
	FuseEnabled            *bool
	HostnameStyle          *string
	IAMRole                *string
	IngressBandwidth       *resource.Quantity
	IMDSRequireToken       *string
	JobAcceptedTimestampMs *uint64
	JobDescriptor          *string
	JobID                  *string
	JobType                *string
	JumboFramesEnabled     *bool
	KvmEnabled             *bool
	LogKeepLocalFile       *bool
	LogUploadCheckInterval *time.Duration
	LogUploadThresholdTime *time.Duration
	LogStdioCheckInterval  *time.Duration
	LogS3WriterIAMRole     *string
	LogS3BucketName        *string
	LogS3PathPrefix        *string
	NetworkBurstingEnabled *bool
	OomScoreAdj            *int32
	PodSchemaVersion       *uint32
	SchedPolicy            *string
	SecurityGroups         *string
	ServiceMeshEnabled     *bool
	ServiceMeshImage       *string
	StaticIPAllocation     *string
	SubnetIDs              *string
	TaskID                 *string
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

	return pConf, err
}
