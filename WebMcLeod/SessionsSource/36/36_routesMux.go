package main

import (
	"io"
	"net/http"
)

type mutexCat int

func (c mutexCat) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Mux cat")
}

type mutexDog int

func (d mutexDog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Mux dog")
}

func main() {
	var dog mutexDog
	var cat mutexCat

	mux := http.NewServeMux()
	mux.Handle("/dog/", dog)
	mux.Handle("/cat/", cat)

	http.ListenAndServe(":8079", mux)

}
