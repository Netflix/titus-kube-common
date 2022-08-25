package resource

const (
	ResourceNameCpu           = "cpu"
	ResourceNameGpu           = "titus/gpu"
	ResourceNameNvidiaGpu     = "nvidia.com/gpu"
	ResourceNameGpuLegacy     = "gpu"
	ResourceNameMemory        = "memory"
	ResourceNameNetwork       = "titus/network"
	ResourceNameNetworkLegacy = "network"
	ResourceNameDisk          = "ephemeral-storage"
	ResourceNameDiskLegacy    = "storage"
	ResourceNamePods          = "pods"
	ResourceNameMockNodesPool = "mock-nodes"
)

func IsMockResourcePool(poolName string) bool {
	return poolName == ResourceNameMockNodesPool
}
