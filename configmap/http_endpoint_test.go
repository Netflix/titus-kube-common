package configmap

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"testing"
)

type testConfig struct {
	StringParameter string `json:"stringParameter"`
}

func TestEncoding(t *testing.T) {
	dc, _ := NewDynamicConfigFromMap(
		map[string]string{"stringParameter": "v"},
		func(rawCurrent map[string]string, previous *ConfigState) (interface{}, error) {
			return testConfig{StringParameter: "v"}, nil
		},
		Options{},
	)
	d := NewDynamicConfigHttpHandler(map[string]DynamicConfig{"test": dc}).(*defaultDynamicConfigHttpHandler)
	var buf bytes.Buffer
	d.writeConfig(&buf, true)
	s := buf.String()
	require.Equal(t, "{\n    \"test\": {\n        \"current\": {\n            \"parsed\": {\n                "+
		"\"stringParameter\": \"v\"\n            },\n            \"raw\": {\n                \"stringParameter\": "+
		"\"v\"\n            }\n        },\n        \"bootstrap\": {\n            \"parsed\": {\n                "+
		"\"stringParameter\": \"v\"\n            },\n            \"raw\": {\n                \"stringParameter\": "+
		"\"v\"\n            }\n        }\n    }\n}\n", s)
}
