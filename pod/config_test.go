package pod

import (
	"testing"
	"time"

	resourceCommon "github.com/Netflix/titus-kube-common/resource"
	"gotest.tools/assert"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	ptr "k8s.io/utils/pointer"
)

func durationPtr(val string) *time.Duration {
	durVal, _ := time.ParseDuration(val)
	return &durVal
}

func resourcePtr(val string) *resource.Quantity {
	resVal, _ := resource.ParseQuantity(val)
	return &resVal
}

func uint32Ptr(val uint32) *uint32 {
	ptrVal := &val
	return ptrVal
}

func uint64Ptr(val uint64) *uint64 {
	ptrVal := &val
	return ptrVal
}

func TestParsePod(t *testing.T) {
	cpu := resource.NewQuantity(1, resource.DecimalSI)
	gpu := resource.NewQuantity(0, resource.DecimalSI)
	mem := resource.NewQuantity(512, resource.BinarySI)
	disk := resource.NewQuantity(999000, resource.BinarySI)
	network := resource.NewQuantity(128, resource.BinarySI)

	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "foo",
			Namespace: "default",
			Annotations: map[string]string{
				// strings
				AnnotationKeyAppDetail:             "mydetail",
				AnnotationKeyAppName:               "myapp",
				AnnotationKeyAppOwnerEmail:         "test@example.com",
				AnnotationKeyAppSequence:           "v000",
				AnnotationKeyAppStack:              "mystack",
				AnnotationKeyIAMRole:               "arn:aws:iam::0:role/DefaultContainerRole",
				AnnotationKeyJobID:                 "myjobid",
				AnnotationKeyJobType:               "BATCH",
				AnnotationKeyJobDescriptor:         "myjobdesc",
				AnnotationKeyPodTitusContainerInfo: "cinfo",

				AnnotationKeyNetworkAccountID:          "123456",
				AnnotationKeyNetworkElasticIPPool:      "pool-1",
				AnnotationKeyNetworkElasticIPs:         "eip-1,eip-2",
				AnnotationKeyNetworkIMDSRequireToken:   "require-token",
				AnnotationKeyNetworkSecurityGroups:     "sg-1,sg-2",
				AnnotationKeyNetworkStaticIPAllocation: "static-ip-alloc",
				AnnotationKeyNetworkSubnetIDs:          "subnet-1,subnet-2",

				// We don't parse these right now - including them so that
				// tests fail if we do start parsing them or remove them
				AnnotationKeyOpportunisticCPU:              "4",
				AnnotationKeyOpportunisticResourceID:       "op-res-id",
				AnnotationKeyPredictionRuntime:             "44",
				AnnotationKeyPredictionConfidence:          "5",
				AnnotationKeyPredictionModelID:             "model-id",
				AnnotationKeyPredictionModelVersion:        "v2",
				AnnotationKeyPredictionABTestCell:          "cell1",
				AnnotationKeyPredictionPredictionAvailable: "a,b",
				AnnotationKeyPredictionSelectorInfo:        "prediction",

				AnnotationKeySecurityAppMetadata:    "app-metadata",
				AnnotationKeySecurityAppMetadataSig: "app-metadata-sig",

				AnnotationKeyPodHostnameStyle: "ec2",
				AnnotationKeyPodSchedPolicy:   "batch",

				AnnotationKeyLogS3BucketName:    "bucket-name",
				AnnotationKeyLogS3PathPrefix:    "s3-prefix",
				AnnotationKeyLogS3WriterIAMRole: "arn:aws:iam::0:role/LogWriterRole",

				AnnotationKeyServiceServiceMeshImage: "titusoss/service-mesh",

				// bools
				AnnotationKeyLogKeepLocalFile:          "true",
				AnnotationKeyNetworkAssignIPv6Address:  "true",
				AnnotationKeyNetworkBurstingEnabled:    "true",
				AnnotationKeyNetworkJumboFramesEnabled: "true",
				AnnotationKeyPodCPUBurstingEnabled:     "true",
				AnnotationKeyPodFuseEnabled:            "true",
				AnnotationKeyPodKvmEnabled:             "true",
				AnnotationKeyServiceServiceMeshEnabled: "true",

				// ints
				AnnotationKeyPodSchemaVersion:       "2",
				AnnotationKeyJobAcceptedTimestampMs: "1602201163007",
				AnnotationKeyPodOomScoreAdj:         "-800",

				// resource values
				AnnotationKeyEgressBandwidth:  "10M",
				AnnotationKeyIngressBandwidth: "20M",

				// durations
				AnnotationKeyLogStdioCheckInterval:  "2m",
				AnnotationKeyLogUploadCheckInterval: "1m",
				AnnotationKeyLogUploadThresholdTime: "3m",
			},
			Labels: map[string]string{
				LabelKeyByteUnitsEnabled: "true",
				LabelKeyCapacityGroup:    "DEFAULT",
				LabelKeyTaskId:           "task-id-in-label",
			},
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{
				{
					Name:  "task-id-in-container",
					Image: "my-registry.example.com/sample/helloworld:latest",
					Resources: corev1.ResourceRequirements{
						Limits: corev1.ResourceList{
							corev1.ResourceCPU:                 *cpu,
							corev1.ResourceMemory:              *mem,
							corev1.ResourceEphemeralStorage:    *disk,
							resourceCommon.ResourceNameGpu:     *gpu,
							resourceCommon.ResourceNameNetwork: *network,
						},
						Requests: corev1.ResourceList{
							corev1.ResourceCPU:                 *cpu,
							corev1.ResourceMemory:              *mem,
							corev1.ResourceEphemeralStorage:    *disk,
							resourceCommon.ResourceNameGpu:     *gpu,
							resourceCommon.ResourceNameNetwork: *network,
						},
					},
				},
			},
		},
	}

	conf, err := PodToConfig(pod)
	assert.NilError(t, err)

	expConf := Config{
		AccountID:              ptr.StringPtr("123456"),
		AppDetail:              ptr.StringPtr("mydetail"),
		AppMetadata:            ptr.StringPtr("app-metadata"),
		AppMetadataSig:         ptr.StringPtr("app-metadata-sig"),
		AppName:                ptr.StringPtr("myapp"),
		AppOwnerEmail:          ptr.StringPtr("test@example.com"),
		AppSequence:            ptr.StringPtr("v000"),
		AppStack:               ptr.StringPtr("mystack"),
		AssignIPv6Address:      ptr.BoolPtr(true),
		BytesEnabled:           ptr.BoolPtr(true),
		CapacityGroup:          ptr.StringPtr("DEFAULT"),
		ContainerInfo:          ptr.StringPtr("cinfo"),
		CPUBurstingEnabled:     ptr.BoolPtr(true),
		EgressBandwidth:        resourcePtr("10M"),
		ElasticIPPool:          ptr.StringPtr("pool-1"),
		ElasticIPs:             ptr.StringPtr("eip-1,eip-2"),
		FuseEnabled:            ptr.BoolPtr(true),
		HostnameStyle:          ptr.StringPtr("ec2"),
		IAMRole:                ptr.StringPtr("arn:aws:iam::0:role/DefaultContainerRole"),
		IMDSRequireToken:       ptr.StringPtr("require-token"),
		IngressBandwidth:       resourcePtr("20M"),
		JobAcceptedTimestampMs: uint64Ptr(1602201163007),
		JobDescriptor:          ptr.StringPtr("myjobdesc"),
		JobID:                  ptr.StringPtr("myjobid"),
		JobType:                ptr.StringPtr("BATCH"),
		JumboFramesEnabled:     ptr.BoolPtr(true),
		KvmEnabled:             ptr.BoolPtr(true),
		LogKeepLocalFile:       ptr.BoolPtr(true),
		LogStdioCheckInterval:  durationPtr("2m"),
		LogUploadCheckInterval: durationPtr("1m"),
		LogUploadThresholdTime: durationPtr("3m"),
		LogS3BucketName:        ptr.StringPtr("bucket-name"),
		LogS3PathPrefix:        ptr.StringPtr("s3-prefix"),
		LogS3WriterIAMRole:     ptr.StringPtr("arn:aws:iam::0:role/LogWriterRole"),
		NetworkBurstingEnabled: ptr.BoolPtr(true),
		OomScoreAdj:            ptr.Int32Ptr(-800),
		PodSchemaVersion:       uint32Ptr(2),
		SchedPolicy:            ptr.StringPtr("batch"),
		SecurityGroups:         ptr.StringPtr("sg-1,sg-2"),
		ServiceMeshEnabled:     ptr.BoolPtr(true),
		ServiceMeshImage:       ptr.StringPtr("titusoss/service-mesh"),
		StaticIPAllocation:     ptr.StringPtr("static-ip-alloc"),
		SubnetIDs:              ptr.StringPtr("subnet-1,subnet-2"),
		TaskID:                 ptr.StringPtr("task-id-in-label"),
	}
	assert.DeepEqual(t, expConf, *conf)
}

// XXX: test all nil
