package configmap

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type callbackState struct {
}

type sampleConfig struct {
	stringParam string
	intParam    int64
}

func sampleConfigMapper(rawCurrent map[string]string, previous *ConfigState) (interface{}, error) {
	config := sampleConfig{
		intParam: ParseInt64(rawCurrent, "intParam", -1),
	}
	if value, ok := rawCurrent["stringParam"]; ok {
		config.stringParam = value
	}
	return &config, nil
}

func TestProcess(t *testing.T) {
	var callbacks []callbackState
	internal := dynamicConfigInternal{
		options: Options{
			OnUpdateCallback: func(current *ConfigState, previous *ConfigState) {
				callbacks = append(callbacks, callbackState{})
			},
		},
		configMapper: sampleConfigMapper,
	}
	err := internal.process(map[string]string{
		"stringParam": "stringValue",
		"intParam":    "123",
	})
	require.Nil(t, err, "unexpected error")
	mapped := internal.current.Mapped.(*sampleConfig)
	require.Equal(t, sampleConfig{stringParam: "stringValue", intParam: 123}, *mapped)
}
