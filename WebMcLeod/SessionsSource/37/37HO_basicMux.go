package main

import (
	"io"
	"net/http"
)

func dog(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "This is the dog router")
}

func me(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Niara the hedgehog")
}

func main() {

	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "This is the root directory of  the exercise \n /dog/, & /me/ routes are available")
	})

	http.ListenAndServe(":8079", nil)
}
