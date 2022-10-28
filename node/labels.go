package node

import (
	corev1 "k8s.io/api/core/v1"
)

const (
	LabelKeyASG               = "node.titus.netflix.com/asg"
	LabelKeyBackend           = "node.titus.netflix.com/backend"
	LabelKeyDecommissioning   = "node.titus.netflix.com/decommissioning"
	LabelKeyInstanceID        = "node.titus.netflix.com/id"
	LabelKeyRemovable         = "node.titus.netflix.com/removable"
	LabelKeyUnremovable       = "node.titus.netflix.com/unremovable"
	LabelKeyResourcePool      = "scaler.titus.netflix.com/resource-pool"
	LabelKeyServerGroupIgnore = "scaler.titus.netflix.com/ignore"
	LabelKeyTerminating       = "node.titus.netflix.com/terminating"
	LabelKeyInstanceType      = "node.kubernetes.io/instance-type"
	LabelKeyMutableBuild      = "node.titus.netflix.com/mutable-build"
	LabelKeyCpuModelName      = "node.titus.netflix.com/cpu-model-name"

	LabelValueBackendMock           = "mock"
	LabelValueBackendVirtualKubelet = "VirtualKubelet"
	LabelValueBackendKubelet        = "kubelet"
)

func IsMockNode(node *corev1.Node) bool {
	if val, ok := node.Labels[LabelKeyBackend]; ok {
		return val == LabelValueBackendMock
	}
	return false
}
