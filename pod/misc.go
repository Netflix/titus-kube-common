package pod

const (
	SchedPriorityFast   = "sched-latency-fast"
	SchedPriorityMedium = "sched-latency-medium"
	SchedPriorityDelay  = "sched-latency-delay"

	SchedNameDefault           = "default-scheduler"
	SchedNameMixed             = "titus-kube-scheduler-mixed"
	SchedNameReserved          = "titus-kube-scheduler-reserved"
	SchedNameRservedBinpacking = "titus-kube-scheduler-reserved-binpacking"
)
