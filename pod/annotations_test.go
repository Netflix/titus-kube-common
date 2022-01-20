package pod

import (
	"testing"

	"github.com/google/go-cmp/cmp/cmpopts"
	"gotest.tools/assert"
	"gotest.tools/assert/cmp"
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

func TestIsPlatformSidecarContainer(t *testing.T) {
	platformSidecarContainer := corev1.Container{Name: "im-a-platform-sidecar"}
	userContainer := corev1.Container{Name: "im-a-user-container"}

	pod := &corev1.Pod{
		ObjectMeta: metav1.ObjectMeta{
			Name:      "foo",
			Namespace: "default",
			Annotations: map[string]string{
				ContainerAnnotation("im-a-platform-sidecar", AnnotationKeySuffixContainersSidecar): "bar",
			},
		},
		Spec: corev1.PodSpec{
			Containers: []corev1.Container{platformSidecarContainer, userContainer},
		},
	}

	assert.Equal(t, IsPlatformSidecarContainer("im-a-platform-sidecar", pod), true)
	assert.Equal(t, IsPlatformSidecarContainer("im-a-user-container", pod), false)
}

func TestPlatformSidecars(t *testing.T) {
	tests := []struct {
		desc         string
		annotations  map[string]string
		wantSidecars []PlatformSidecar
	}{
		{
			desc: "no platform sidecars",
			annotations: map[string]string{
				"pod.netflix.com/cpu-bursting-enabled": "true",
			},
			wantSidecars: nil,
		},
		{
			desc: "one platform sidecar",
			annotations: map[string]string{
				"healthz-v1.platform-sidecars.netflix.com":         "true",
				"healthz-v1.platform-sidecars.netflix.com/channel": "stable",
			},
			wantSidecars: []PlatformSidecar{
				{
					Name:    "healthz-v1",
					Channel: "stable",
				},
			},
		},
		{
			desc: "one platform sidecar with arguments",
			annotations: map[string]string{
				"healthz-v2.platform-sidecars.netflix.com":           "true",
				"healthz-v2.platform-sidecars.netflix.com/channel":   "staging",
				"healthz-v2.platform-sidecars.netflix.com/arguments": `{"cmd": "echo", "cmdArgs": "healthy"}`,
			},
			wantSidecars: []PlatformSidecar{
				{
					Name:     "healthz-v2",
					Channel:  "staging",
					ArgsJSON: []byte(`{"cmd": "echo", "cmdArgs": "healthy"}`),
				},
			},
		},
		{
			desc: "one platform sidecar set to false",
			annotations: map[string]string{
				"healthz-v1.platform-sidecars.netflix.com":         "false",
				"healthz-v1.platform-sidecars.netflix.com/channel": "stable",
			},
			wantSidecars: nil,
		},
		{
			desc: "many platform sidecars",
			annotations: map[string]string{
				"healthz-v1.platform-sidecars.netflix.com":           "false",
				"healthz-v1.platform-sidecars.netflix.com/channel":   "stable",
				"pod.netflix.com/cpu-bursting-enabled":               "true",
				"healthz-v2.platform-sidecars.netflix.com":           "true",
				"healthz-v2.platform-sidecars.netflix.com/channel":   "staging",
				"healthz-v2.platform-sidecars.netflix.com/arguments": `{"cmd": "echo", "cmdArgs": "healthy"}`,
				"healthz-v3.platform-sidecars.netflix.com":           "true",
				"healthz-v3.platform-sidecars.netflix.com/channel":   "stable",
			},
			wantSidecars: []PlatformSidecar{
				{
					Name:     "healthz-v2",
					Channel:  "staging",
					ArgsJSON: []byte(`{"cmd": "echo", "cmdArgs": "healthy"}`),
				},
				{
					Name:    "healthz-v3",
					Channel: "stable",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			sidecars, err := PlatformSidecars(tt.annotations)
			assert.NilError(t, err, "PlatformSidecars(%+v)", tt.annotations)

			lessThan := func(a, b PlatformSidecar) bool {
				return a.Name < b.Name
			}
			assert.Check(t, cmp.DeepEqual(tt.wantSidecars, sidecars, cmpopts.SortSlices(lessThan)), "PlatformSidecars(%+v)", tt.annotations)
		})
	}
}

func TestPlatformSidecarsErrors(t *testing.T) {
	tests := []struct {
		desc            string
		annotations     map[string]string
		wantErrContains string
	}{
		{
			desc: "sidecar include annotation is not a bool",
			annotations: map[string]string{
				"healthz-v1.platform-sidecars.netflix.com": "maybe",
			},
			wantErrContains: `sidecar annotation "healthz-v1.platform-sidecars.netflix.com" must be a bool value`,
		},
		{
			desc: "sidecar is missing channel annotation",
			annotations: map[string]string{
				"healthz-v1.platform-sidecars.netflix.com":           "true",
				"healthz-v1.platform-sidecars.netflix.com/arguments": `{"cmd": "echo", "cmdArgs": "healthy"}`,
			},
			wantErrContains: `sidecar "healthz-v1.platform-sidecars.netflix.com" must have a channel specified via annotation "healthz-v1.platform-sidecars.netflix.com/channel"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.desc, func(t *testing.T) {
			_, err := PlatformSidecars(tt.annotations)
			assert.Check(t, cmp.ErrorContains(err, tt.wantErrContains), "PlatformSidecars(%+v)", tt.annotations)
		})
	}
}
