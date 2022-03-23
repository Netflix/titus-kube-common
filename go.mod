module github.com/Netflix/titus-kube-common

go 1.15

require (
	github.com/docker/distribution v2.7.1+incompatible
	github.com/go-logr/logr v0.3.0
	github.com/google/go-cmp v0.5.5
	github.com/hashicorp/go-multierror v1.0.0
	github.com/kr/text v0.2.0 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/stretchr/testify v1.6.1
	golang.org/x/sys v0.0.0-20210119212857-b64e53b001e4 // indirect
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gotest.tools v2.2.0+incompatible
	k8s.io/api v0.20.9
	k8s.io/apimachinery v0.20.9
	k8s.io/client-go v0.20.9
	k8s.io/utils v0.0.0-20210111153108-fddb29f9d009
	sigs.k8s.io/controller-runtime v0.8.3
)

replace (
	k8s.io/api => k8s.io/api v0.20.9
	k8s.io/apiextensions-apiserver => k8s.io/apiextensions-apiserver v0.20.9
	k8s.io/apimachinery => k8s.io/apimachinery v0.20.9
	k8s.io/apiserver => k8s.io/apiserver v0.20.9
	k8s.io/cli-runtime => k8s.io/cli-runtime v0.20.9
	k8s.io/client-go => k8s.io/client-go v0.20.9
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.20.9
	k8s.io/cluster-bootstrap => k8s.io/cluster-bootstrap v0.20.9
	k8s.io/code-generator => k8s.io/code-generator v0.20.9
	k8s.io/component-base => k8s.io/component-base v0.20.9
	k8s.io/cri-api => k8s.io/cri-api v0.20.9
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.20.9
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.20.9
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.20.9
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.20.9
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.20.9
	k8s.io/kubectl => k8s.io/kubectl v0.20.9
	k8s.io/kubelet => k8s.io/kubelet v0.20.9
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.20.9
	k8s.io/metrics => k8s.io/metrics v0.20.9
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.20.9
)
