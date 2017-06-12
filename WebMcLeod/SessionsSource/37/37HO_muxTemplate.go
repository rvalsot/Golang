package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

func main() {
	http.Handle("/fifufu", http.HandlerFunc(fifufu))
	http.HandleFunc("/bororo", bororo)
	http.HandleFunc("/", templade)

	http.ListenAndServe(":8079", nil)
}

// -----------------------------------------------------------------------

func fifufu(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Fifufu endpoint")
}

func bororo(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Bororo endpoint")
}

func templade(w http.ResponseWriter, r *http.Request) {

	templ, err := template.ParseFiles("template1.gohtml")
	if err != nil {
		log.Fatalln("Error parsing file: \n", err)
	}

	dataMap := "This is a data map to be deployed in this ListenAndServe Mux template."

	err = templ.ExecuteTemplate(w, "template1.gohtml", dataMap)
	if err != nil {
		log.Fatalln("Error while executing template: \n", err)
	}

}
