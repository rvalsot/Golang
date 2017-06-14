package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var templat *template.Template

func init() {
	var err error
	templat, err = template.ParseFiles("templates/index.gohtml")
	if err != nil {
		panic(err)
	}
}

func index(w http.ResponseWriter, r *http.Request) {
	templat.Execute(w, nil)
}

func main() {
	http.HandleFunc("/", index)

	port := ":8070"
	fmt.Println("If no error, listening and serving at port", port)

	http.ListenAndServe(port, nil)
}
