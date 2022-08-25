package resourcepool

const (
	ResourcePoolElastic   = "elastic"
	ResourcePoolReserved  = "reserved"
	ResourcePoolMockNodes = "mock-nodes"
)

func IsMockResourcePool(poolName string) bool {
	return poolName == ResourcePoolMockNodes
}
