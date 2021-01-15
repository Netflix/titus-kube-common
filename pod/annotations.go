package pod

const (
	AnnotationKeyInstanceType = "node.titus.netflix.com/itype"
	AnnotationKeyRegion       = "node.titus.netflix.com/region"
	AnnotationKeyZone         = "node.titus.netflix.com/zone"

	// Pod Networking
	AnnotationKeyEgressBandwidth  = "kubernetes.io/egress-bandwidth"
	AnnotationKeyIngressBandwidth = "kubernetes.io/ingress-bandwidth"
	AnnotationKeySecurityGroups   = "network.titus.netflix.com/securityGroups"
	AnnotationKeySubnets          = "network.titus.netflix.com/subnets"
	AnnotationKeyAccountID        = "network.titus.netflix.com/accountId"

	// Security
	AnnotationKeyIAMRole = "iam.amazonaws.com/role"

	// Pod ENI
	AnnotationKeyIPv4Address      = "network.titus.netflix.com/address-ipv4"
	AnnotationKeyIPv4PrefixLength = "network.titus.netflix.com/prefixlen-ipv4"
	AnnotationKeyIPv6Address      = "network.titus.netflix.com/address-ipv6"
	AnnotationKeyIPv6PrefixLength = "network.titus.netflix.com/prefixlen-ipv6"

	AnnotationKeyBranchEniID     = "network.titus.netflix.com/branch-eni-id"
	AnnotationKeyBranchEniMac    = "network.titus.netflix.com/branch-eni-mac"
	AnnotationKeyBranchEniVpcID  = "network.titus.netflix.com/branch-eni-vpc"
	AnnotationKeyBranchEniSubnet = "network.titus.netflix.com/branch-eni-subnet"

	AnnotationKeyTrunkEniID    = "network.titus.netflix.com/trunk-eni-id"
	AnnotationKeyTrunkEniMac   = "network.titus.netflix.com/trunk-eni-mac"
	AnnotationKeyTrunkEniVpcID = "network.titus.netflix.com/trunk-eni-vpc"

	AnnotationKeyVlanID        = "network.titus.netflix.com/vlan-id"
	AnnotationKeyAllocationIdx = "network.titus.netflix.com/allocation-idx"
)
