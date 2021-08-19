package urlshort

import (
	"gopkg.in/yaml.v2"
	"net/http"
)

type redirectItem struct {
	Path string `yaml:"path"`
	URL string `yaml:"url"`
}

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

func YAMLHandler(yaml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	parsedYaml, err := parseYAML(yaml)
	if err != nil {
		return fallback.ServeHTTP, err
	}
	pathMap := buildMap(parsedYaml)
	return MapHandler(pathMap, fallback), nil
}

func parseYAML(yamlBytes []byte) ([]redirectItem, error) {
	var yamlList []redirectItem
	err := yaml.Unmarshal(yamlBytes, &yamlList)
	if err != nil {
		return nil, err
	}
	return yamlList, nil
}

func buildMap(yamlList []redirectItem) map[string]string {
	var yamlMap = map[string]string{}
	var entry redirectItem
		for _, entry = range yamlList {
			yamlMap[entry.Path] = entry.URL
		}
	return yamlMap
}
