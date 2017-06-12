package main

import (
	"io"
	"net/http"
)

type webDog int

func (dog webDog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/dog":
		io.WriteString(w, "Dog path")
	case "/cat":
		io.WriteString(w, "Cat path")
	}
}

func main() {
	var doo webDog
	http.ListenAndServe(":8079", doo)
}
