package main

import (
	"fmt"
	"net/http"
)

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	fmt.Fprintln(w, "Go to look at your terminal")
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	port := ":8070"

	fmt.Println("To ListenAndServe at port", port)
	http.ListenAndServe(port, nil)
}
