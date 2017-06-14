package main

import (
	"fmt"
	"io"
	"net/http"
)

func hedgehog4(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="niara.jpeg"`)
}

func hedgehogPic4(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "niara.jpeg")
}

func main() {
	port := ":8070"

	http.HandleFunc("/", hedgehog4)
	http.HandleFunc("/niara.jpeg", hedgehogPic4)
	http.ListenAndServe(port, nil)
	fmt.Println("Listening and Serving at localhost", port)
}
