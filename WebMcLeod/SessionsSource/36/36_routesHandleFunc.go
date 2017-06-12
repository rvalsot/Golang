package main

import (
	"io"
	"net/http"
)

func dog(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Dooog mux function")
}

func main() {
	http.HandleFunc("/dog/", dog)

	http.HandleFunc("/cow", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Cow-in-main mux function")
	})

	http.ListenAndServe(":8079", nil)
}
