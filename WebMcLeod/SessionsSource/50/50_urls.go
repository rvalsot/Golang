package main

import (
	"fmt"
	"net/http"
)

func fifufu(w http.ResponseWriter, r *http.Request) {
	val := r.FormValue("q")
	fmt.Fprintln(w, "Do my search:"+val)
}

func main() {
	http.HandleFunc("/", fifufu)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	port := ":8070"

	fmt.Println("To be Listened And Served at port", port)
	http.ListenAndServe(port, nil)
}
