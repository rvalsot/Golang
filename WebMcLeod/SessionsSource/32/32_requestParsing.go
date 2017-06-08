package main

import (
	"html/template"
	"log"
	"net/http"
)

type webHotdog int

var templat *template.Template

func (m webHotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	if err != nil {
		log.Fatalln(err)
	}

	templat.ExecuteTemplate(w, "index1.gohtml", req.Form)
}

func init() {
	templat = template.Must(template.ParseFiles("index1.gohtml"))
}

func main() {
	var d webHotdog

	http.ListenAndServe(":8080", d)
}
