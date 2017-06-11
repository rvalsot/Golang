package main

import (
	"html/template"
	"log"
	"net/http"
	"net/url"
)

type webSunday int

func (s webSunday) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	data := struct {
		Method      string
		Submissions url.Values
	}{
		r.Method,
		r.Form,
	}
	templa.ExecuteTemplate(w, "index2.gohtml", data)
}

var templa *template.Template

func init() {
	templa = template.Must(template.ParseFiles("index2.gohtml"))
}

func main() {
	var d webSunday
	http.ListenAndServe(":8081", d)
}
