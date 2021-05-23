package urlshort

import (
	"net/http"

	"gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {

		//getting the path from the request
		path := r.URL.Path
		//if we match the path then we redirect
		if destUrl, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, destUrl, http.StatusFound)
			return
		}
		//calling fallback if we do not get the url for the given path
		fallback.ServeHTTP(w, r)

	}
}

type pathURL struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//     - path: /some-path
//       url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yamlBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {
	//1. Parse the yaml
	pathURLs, err := parseYAML(yamlBytes)
	if err != nil {
		return nil, err
	}
	//2. Convert yaml array into map
	pathsToUrls := buildMap(pathURLs)
	//3. return the MapHandller using map
	return MapHandler(pathsToUrls, fallback), nil
}

//parseYAML parses a byte slice yaml into a list of pathURL type
func parseYAML(data []byte) ([]pathURL, error) {
	var pathURLs []pathURL
	err := yaml.Unmarshal(data, &pathURLs)
	if err != nil {
		return nil, err
	}
	return pathURLs, nil
}

//buildMap builds a map out of []pathURL
func buildMap(list []pathURL) map[string]string {
	pathsToUrls := make(map[string]string)
	for _, pu := range list {
		pathsToUrls[pu.Path] = pu.URL
	}
	return pathsToUrls
}
