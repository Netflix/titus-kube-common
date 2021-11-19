package pod

import (
	corev1 "k8s.io/api/core/v1"
)

func GetUserContainer(pod *corev1.Pod) *corev1.Container {
	if len(pod.Spec.Containers) == 0 {
		return nil
	}

	firstContainer := pod.Spec.Containers[0]
	for i := range pod.Spec.Containers {
		c := &pod.Spec.Containers[i]
		if c.Name == pod.Name {
			return c
		}
	}

	return &firstContainer
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
