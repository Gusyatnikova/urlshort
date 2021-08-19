package urlshort

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
	"net/http"
)

type redirectItemYAML struct {
	Path string `yaml:"path"`
	URL string `yaml:"url"`
}

type redirectItemJSON map[string]string

func MapHandler(inToRedirected map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		val, ok := inToRedirected[r.URL.String()]
		if ok {
			http.Redirect(w, r, val, http.StatusMovedPermanently)
		} else {
			fallback.ServeHTTP(w, r)
		}
	}
}

func JSONHandler(json []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedJson, err := parseJSON(json)
	if err != nil {
		return fallback.ServeHTTP, err
	}
	return MapHandler(parsedJson, fallback), nil
}

func YAMLHandler(yaml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYaml, err := parseYAML(yaml)
	if err != nil {
		return fallback.ServeHTTP, err
	}
	pathMap := buildMap(parsedYaml)
	return MapHandler(pathMap, fallback), nil
}

func parseYAML(yamlBytes []byte) ([]redirectItemYAML, error) {
	var yamlList []redirectItemYAML
	err := yaml.Unmarshal(yamlBytes, &yamlList)
	if err != nil {
		return nil, err
	}
	return yamlList, nil
}

func parseJSON(jsonBytes []byte) (redirectItemJSON, error) {
	var jsonList redirectItemJSON
	err := json.Unmarshal(jsonBytes, &jsonList)
	if err != nil {
		return nil, err
	}
	return jsonList, nil
}

func buildMap(yamlList []redirectItemYAML) map[string]string {
	var yamlMap = map[string]string{}
	var entry redirectItemYAML
		for _, entry = range yamlList {
			yamlMap[entry.Path] = entry.URL
		}
	return yamlMap
}
