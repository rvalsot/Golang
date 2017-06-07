package main

import (
	"fmt"
	"net/http"
)

type hotdog int

func (mustard hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Println(w, "This is a message")
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
