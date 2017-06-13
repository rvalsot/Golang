package main

import (
	"io"
	"net/http"
	"os"
)

func hedgehog(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `
    <img src="/niara.jpeg" />
  `)
}

func hedgehogPic(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("niara.jpeg")
	if err != nil {
		http.Error(w, "file not found", 404)
		return
	}
	defer file.Close()

	io.Copy(w, file)
}

func main() {
	http.HandleFunc("/", hedgehog)
	http.HandleFunc("/niara", hedgehogPic)
	http.ListenAndServe(":8070", nil)
}
