package configmap

import (
	"io"
	"net/http"
	"strconv"

	"k8s.io/apimachinery/pkg/util/json"
)

type defaultDynamicConfigHttpHandler struct {
	configMap map[string]DynamicConfig
}

type configSection struct {
	Parsed interface{}       `json:"parsed"`
	Raw    map[string]string `json:"raw"`
}

type configRepresentation struct {
	Current   *configSection `json:"current"`
	Bootstrap *configSection `json:"bootstrap,omitempty"`
}

func NewDynamicConfigHttpHandler(configMap map[string]DynamicConfig) http.Handler {
	return &defaultDynamicConfigHttpHandler{
		configMap: configMap,
	}
}

func (d *defaultDynamicConfigHttpHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	if request.Method != "GET" {
		response.WriteHeader(405)
		return
	}

	all := false
	if v, ok := request.URL.Query()["all"]; ok {
		if len(v) == 0 || len(v[0]) == 0 {
			all = true
		} else {
			all, _ = strconv.ParseBool(v[0])
		}
	}
	response.WriteHeader(200)
	response.Header().Set("Content-Type", "application/json")
	d.writeConfig(response, all)
}

func (d *defaultDynamicConfigHttpHandler) writeConfig(writer io.Writer, all bool) {
	result := map[string]configRepresentation{}
	for name, config := range d.configMap {
		representation := configRepresentation{
			Current: &configSection{
				Raw:    config.GetRaw(),
				Parsed: config.Get(),
			},
		}
		if all {
			representation.Bootstrap = &configSection{
				Raw:    config.GetRawBootstrap(),
				Parsed: config.GetBootstrap(),
			}
		}
		result[name] = representation
	}

	encoder := json.NewEncoder(writer)
	encoder.SetIndent("", "    ")
	_ = encoder.Encode(result)
}
