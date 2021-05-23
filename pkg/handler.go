package urlshort

import (
	"net/http"
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
