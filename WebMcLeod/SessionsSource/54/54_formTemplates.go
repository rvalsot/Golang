package main

import (
	"html/template"
	"log"
	"net/http"
)

var templat *template.Template

func init() {
	templat = template.Must(template.ParseGlob("templates/*"))
}

type person struct {
	FirstName  string
	LastName   string
	Subscribed bool
}

func formTemplates(w http.ResponseWriter, r *http.Request) {
	first := r.FormValue("first")
	last := r.FormValue("last")
	subs := r.FormValue("subscribe") == "on"

	err := templat.ExecuteTemplate(w, "index.gohtml", person{first, last, subs})
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}
}

func main() {
	http.HandleFunc("/", formTemplates)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	port := ":8070"
	http.ListenAndServe(port, nil)
}
