module github.com/Netflix/titus-kube-common

go 1.13

replace (
	k8s.io/api => k8s.io/kubernetes/staging/src/k8s.io/api v0.0.0-20200118001809-59603c6e503c
	k8s.io/apimachinery => k8s.io/kubernetes/staging/src/k8s.io/apimachinery v0.0.0-20200118001809-59603c6e503c
)

require (
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/google/go-cmp v0.5.0 // indirect
	github.com/hashicorp/go-multierror v1.0.0
	github.com/kr/text v0.2.0 // indirect
	github.com/niemeyer/pretty v0.0.0-20200227124842-a10e7caefd8e // indirect
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/stretchr/testify v1.6.1 // indirect
	golang.org/x/net v0.0.0-20200625001655-4c5254603344 // indirect
	gopkg.in/check.v1 v1.0.0-20200227125254-8fa46927fb4f // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
	gotest.tools v2.2.0+incompatible
	k8s.io/api v0.17.2
	k8s.io/apimachinery v0.17.2
	k8s.io/utils v0.0.0-20201104234853-8146046b121e
)
