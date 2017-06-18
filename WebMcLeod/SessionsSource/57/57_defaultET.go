package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templ *template.Template

func init() {
	templ = template.Must(template.ParseGlob("defaultT/*"))
}

type person struct {
	FirstName string
	LastName  string
	Suscribed bool
}

func fo(w http.ResponseWriter, r *http.Request) {
	bs := make([]byte, r.ContentLength)

	r.Body.Read(bs)
	body := string(bs)

	err := templ.ExecuteTemplate(w, "index.gohtml", body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}
}

func main() {
	http.HandleFunc("/", fo)
	http.Handle("/favicon.icon", http.NotFoundHandler())

	port := ":8070"
	fmt.Println("Listening and Serving at port", port)
	http.ListenAndServe(port, nil)
}
