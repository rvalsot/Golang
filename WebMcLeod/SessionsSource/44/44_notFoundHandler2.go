package main

import (
	"fmt"
	"net/http"
)

func fioo(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	fmt.Println(r.URL.Path)
	fmt.Fprintln(w, "Go to look at your terminal")
}

func main() {
	http.HandleFunc("/", fioo)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	port := ":8070"
	fmt.Println("To be Listened and Served at port", port)
	http.ListenAndServe(port, nil)

}
