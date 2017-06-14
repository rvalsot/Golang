package main

import (
	"fmt"
	"io"
	"net/http"
)

func hedgehog2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/resources/niara.jpeg" />`)
}

func main() {
	http.HandleFunc("/", hedgehog2)
	http.Handle("/resources/",
		http.StripPrefix("/resources",
			http.FileServer(
				http.Dir("./photos"))))

	port := ":8070"
	fmt.Println("Listening and Serving from port ", port)

	http.ListenAndServe(port, nil)
}
