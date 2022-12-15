package node

const (
	AnnotationKeyAccount        = "node.titus.netflix.com/account"
	AnnotationKeyAccountID      = "node.titus.netflix.com/accountId"
	AnnotationKeyAMI            = "node.titus.netflix.com/ami"
	AnnotationKeyASG            = "node.titus.netflix.com/asg"
	AnnotationKeyCluster        = "node.titus.netflix.com/cluster"
	AnnotationKeyENIResourceSet = "node.titus.netflix.com/res"
	AnnotationKeyInstanceID     = "node.titus.netflix.com/id"
	AnnotationKeyInstanceType   = "node.titus.netflix.com/itype"
	AnnotationKeyRegion         = "node.titus.netflix.com/region"
	AnnotationKeyZone           = "node.titus.netflix.com/zone"
	AnnotationKeyStack          = "node.titus.netflix.com/stack"
	// AnnotationKeyNodeTerminationReason is a human readable string indicating *why* a node was terminated.
	// It is the responsibility of any code that calls node.delete() to populate this annotation so that operators
	// of the system can understand why a node was deleted.
	AnnotationKeyNodeTerminationReason = "node.titus.netflix.com/node-termination-reason"
	// AnnotationKeyNodeTerminationByCaller is a human readable string indicating which Titus component actually
	// deleted the a node, to aid operators in investigating "why did this node go away".
	AnnotationKeyNodeTerminationByCaller = "node.titus.netflix.com/node-termination-by-caller"
)
