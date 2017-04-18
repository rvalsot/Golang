package main

import (
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {
	// http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
	// 	w.Write([]byte("Niara is a Hedgehog"))
	// })
	http.Handle("/", new(MyHandler))

	http.ListenAndServe(":8000", nil)
}

// MyHandler is a personal handler
type MyHandler struct {
	http.Handler
}

func (thisHandler *MyHandler) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	path := "files" + req.URL.Path
	data, err := ioutil.ReadFile(string(path))

	if err == nil {
		var contentType string
		w.Write(data)
		if strings.HasSuffix(path, ".css") {
			contentType = "text/css"
		} else if strings.HasSuffix(path, ".html") {
			contentType = "text/html"
		} else if strings.HasSuffix(path, ".js") {
			contentType = "application/javascript"
		} else if strings.HasSuffix(path, ".png") {
			contentType = "image/png"
		} else {
			contentType = "text/plain"
		}
	} else {
		w.WriteHeader(404)
		w.Write([]byte("Error 404 in your action"))
	}
}
