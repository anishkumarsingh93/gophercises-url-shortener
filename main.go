package main

import (
	"fmt"
	"net/http"

	urlshort "github.com/anishkumarsingh93/gophercises-url-shortener/pkg"
)

var port = ":8080"

func main() {

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
		"/google":         "https://google.com",
	}

	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	//starting the server
	fmt.Println("Starting server at port: ", port)
	http.ListenAndServe(port, mapHandler)

}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World!")
}
