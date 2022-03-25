package pod

import (
	corev1 "k8s.io/api/core/v1"
)

func GetMainUserContainer(pod *corev1.Pod) *corev1.Container {
	if len(pod.Spec.Containers) == 0 {
		return nil
	}

	// Older method where the main container's name was the taskid
	for i := range pod.Spec.Containers {
		c := &pod.Spec.Containers[i]
		if c.Name == pod.Name {
			return c
		}
	}

	// Newer method where the main container's name is "main"
	for i := range pod.Spec.Containers {
		c := &pod.Spec.Containers[i]
		if c.Name == "main" {
			return c
		}
	}

	// Fallback method, whatever came first
	return &pod.Spec.Containers[0]
}

func GetContainerByName(pod *corev1.Pod, name string) *corev1.Container {
	for i := range pod.Spec.Containers {
		c := &pod.Spec.Containers[i]
		if c.Name == name {
			return c
		}
	}

	return nil
}

// GetImageTagForContainer looks up the original tag that was used to create
// the image string in the Container Spec.
// It may return an empty string if there was no tag, or if it was missing
func GetImageTagForContainer(cName string, pod *corev1.Pod) string {
	key := AnnotationKeyImageTagPrefix + cName
	value := pod.ObjectMeta.Annotations[key]
	return value
}
