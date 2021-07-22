package pod

import (
	"testing"

	"gotest.tools/assert"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestPodTaskID(t *testing.T) {
	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "foo",
			Namespace: "default",
			Labels: map[string]string{
				LabelKeyTaskId: "00000000-0000-0000-0000-000000000000",
			},
		},
	}

	id, err := TaskID(pod)
	assert.NilError(t, err)
	assert.Equal(t, id, "00000000-0000-0000-0000-000000000000")

	unlabeledPod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "foo",
			Namespace: "default",
		},
	}

	id, err = TaskID(unlabeledPod)
	assert.Error(t, err, "pod doesn't contain label "+LabelKeyTaskId)
}

func TestPodJobID(t *testing.T) {
	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "foo",
			Namespace: "default",
			Labels: map[string]string{
				LabelKeyJobId: "00000000-0000-0000-0000-000000000000",
			},
		},
	}

	id, err := JobID(pod)
	assert.NilError(t, err)
	assert.Equal(t, id, "00000000-0000-0000-0000-000000000000")

	unlabeledPod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "foo",
			Namespace: "default",
		},
	}

	id, err = JobID(unlabeledPod)
	assert.Error(t, err, "pod doesn't contain label "+LabelKeyJobId)
}
