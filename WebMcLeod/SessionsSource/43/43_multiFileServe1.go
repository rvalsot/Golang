package main

import (
	"fmt"
	"io"
	"net/http"
)

func hedgehog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="niara.jpeg" />`)
}

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc("/niara", hedgehog)

	port := ":8070"
	fmt.Println("Going to listen and serve at port ", port)

	http.ListenAndServe(port, nil)
}
