package main

import (
	"html/template"
	"net/http"
)

var templat *template.Template

func init() {
	templat = template.Must(template.ParseGlob("templates/*"))
}

func index(w http.ResponseWriter, r *http.Request) {
	templat.ExecuteTemplate(w, "index.gohtml", nil)
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	port := ":8070"
	http.ListenAndServe(port, nil)
}
