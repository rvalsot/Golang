package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

var templa *template.Template

// type person struct {
// 	FirstName  string
// 	LastName   string
// 	Susbcribed bool
// }

func init() {
	templa = template.Must(template.ParseGlob("multipartT/*"))
}

func main() {
	http.HandleFunc("/", foo)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	port := ":8070"
	fmt.Println("Listening and Serving at port", port)

	http.ListenAndServe(port, nil)
}

func foo(w http.ResponseWriter, r *http.Request) {
	bs := make([]byte, r.ContentLength)

	r.Body.Read(bs)
	body := string(bs)

	err := templa.ExecuteTemplate(w, "index.gohtml", body)
	if err != nil {
		http.Error(w, err.Error(), 500)
		log.Fatalln(err)
	}

}
