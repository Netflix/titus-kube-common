package pod

import (
	corev1 "k8s.io/api/core/v1"
)

const (
	// High-level "domains" used for labels and annotations
	DomainNetflix = "netflix.com"
	DomainTitus   = "titus.netflix.com"
	DomainPod     = "pod.netflix.com"

	// Job details
	LabelKeyAppLegacy      = "netflix.com/applicationName"
	LabelKeyDetailLegacy   = "netflix.com/detail"
	LabelKeySequenceLegacy = "netflix.com/sequence"
	LabelKeyStackLegacy    = "netflix.com/stack"

	LabelKeyByteUnitsEnabled    = "pod.titus.netflix.com/byteUnits"
	LabelKeyCapacityGroupLegacy = "titus.netflix.com/capacityGroup"

	// v1 pod labels
	LabelKeyJobId            = "v3.job.titus.netflix.com/job-id"
	LabelKeyTaskId           = "v3.job.titus.netflix.com/task-id"
	LabelKeyCapacityGroup    = "titus.netflix.com/capacity-group"
	LabelKeyWorkloadName     = "workload.netflix.com/name"
	LabelKeyWorkloadStack    = "workload.netflix.com/stack"
	LabelKeyWorkloadDetail   = "workload.netflix.com/detail"
	LabelKeyWorkloadSequence = "workload.netflix.com/sequence"
)

func parseLabels(pod *corev1.Pod, pConf *Config) error {
	labels := pod.GetLabels()

	// Only parse the labels that aren't duplicates of annotations
	cVal, ok := labels[LabelKeyCapacityGroup]
	if ok {
		pConf.CapacityGroup = &cVal
	}

	// Maybe pull this from the containers in the pod instead?
	tVal, ok := labels[LabelKeyTaskId]
	if ok {
		pConf.TaskID = &tVal
	}

	return nil
}
