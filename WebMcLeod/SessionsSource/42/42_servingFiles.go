package main

import (
	"io"
	"net/http"
)

func main() {
	http.HandleFunc("/", puppy)
	http.ListenAndServe(":8070", nil)
}

func puppy(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
    <!-- not from our localhost, but from wikipedia -->
    <img src="https://upload.wikimedia.org/wikipedia/commons/a/a9/Hedgehog_cyprus_hg.jpg" />
  `)
}
