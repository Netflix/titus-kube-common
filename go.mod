module github.com/Netflix/titus-kube-common

go 1.15

replace (
	k8s.io/api => k8s.io/api v0.19.10
	k8s.io/apimachinery => k8s.io/apimachinery v0.19.10
)

require (
	github.com/docker/distribution v2.7.1+incompatible
	github.com/google/go-cmp v0.5.0 // indirect
	github.com/hashicorp/go-multierror v1.0.0
	github.com/kr/text v0.2.0 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/stretchr/testify v1.6.1 // indirect
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
	gotest.tools v2.2.0+incompatible
	k8s.io/api v0.19.10
	k8s.io/apimachinery v0.19.10
	k8s.io/client-go v0.19.10
	k8s.io/klog/v2 v2.2.0 // indirect
	k8s.io/utils v0.0.0-20200912215256-4140de9c8800
)
