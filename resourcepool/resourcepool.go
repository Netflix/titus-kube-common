package resourcepool

const (
	ResourcePoolMockPods = "mock-nodes"
)

func IsMockResourcePool(poolName string) bool {
	return poolName == ResourcePoolMockPods
}
