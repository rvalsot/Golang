package main

import (
	"html/template"
	"log"
	"net/http"
)

var templat *template.Template

type person struct {
	FirstName  string
	LastName   string
	Subscribed string
}

func init() {
	templat = template.Must(template.ParseGlob("textT/*"))
}

func fooo(w http.ResponseWriter, r *http.Request) {
	bs := make([]byte, r.ContentLength)
	r.Body.Read(bs)
	body := string(bs)

	err := templat.ExecuteTemplate(w, "index.gohtml", body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}
}

func main() {
	http.HandleFunc("/", fooo)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	port := ":8070"

	http.ListenAndServe(port, nil)

}
