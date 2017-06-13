package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func hedgehog3(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	io.WriteString(w, `<img src="/niara.jpeg"/>`)
}

func hedgehogPic3(w http.ResponseWriter, r *http.Request) {
	file, err := os.Open("niara.jpeg")
	if err != nil {
		http.Error(w, "File not found", 404)
		return
	}

	defer file.Close()

	fi, err := file.Stat()
	if err != nil {
		http.Error(w, "File not found", 404)
		return
	}
	fmt.Fprintf(w, "Filename: %s", file.Name())
	http.ServeContent(w, r, file.Name(), fi.ModTime(), file)

}

func main() {
	http.HandleFunc("/", hedgehog3)
	http.HandleFunc("/niara.jpeg", hedgehog3)

	http.ListenAndServe(":8070", nil)
}
