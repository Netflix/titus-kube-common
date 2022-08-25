package resourcepool

const (
	ResourcePoolElastic    = "elastic"
	ResourcePoolElasticGPU = "elasticGpu"
	ResourcePoolReserved   = "reserved"
	ResourcePoolMocNodes   = "mock-nodes"
)

func IsMockResourcePool(poolName string) bool {
	return poolName == ResourcePoolMocNodes
}
