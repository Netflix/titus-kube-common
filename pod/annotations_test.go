package pod

import (
	"testing"

	"gotest.tools/assert"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestPodSchemaVersion(t *testing.T) {
	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "foo",
			Namespace: "default",
			Annotations: map[string]string{
				AnnotationKeyPodSchemaVersion: "5",
			},
		},
	}

	ver, err := PodSchemaVersion(pod)
	assert.NilError(t, err)
	assert.Equal(t, uint32(5), ver)
}

func TestPodSchemaVersionUnset(t *testing.T) {
	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "foo",
			Namespace: "default",
		},
	}

	ver, err := PodSchemaVersion(pod)
	assert.NilError(t, err)
	assert.Equal(t, uint32(0), ver)
}

func TestBadPodSchemaVersion(t *testing.T) {
	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "foo",
			Namespace: "default",
			Annotations: map[string]string{
				AnnotationKeyPodSchemaVersion: "asdf",
			},
		},
	}

	_, err := PodSchemaVersion(pod)
	assert.ErrorContains(t, err, "annotation is not a valid uint32 value: "+AnnotationKeyPodSchemaVersion)
}

func TestPodPlatformSidecars(t *testing.T) {
	platformSidecarContainer := corev1.Container{Name: "im-a-platform-sidecar"}
	userContainer := corev1.Container{Name: "im-a-user-container"}

	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "foo",
			Namespace: "default",
			Annotations: map[string]string{
				AnnotationKeyPrefixContainerType + "im-a-platform-sidecar": AnnotationValueContainerTypePlatformSidecar,
			},
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{platformSidecarContainer, userContainer},
		},
	}

	assert.Equal(t, IsPlatformSidecarContainer("im-a-platform-sidecar", pod), true)
	assert.Equal(t, IsPlatformSidecarContainer("im-a-user-container", pod), false)
}
