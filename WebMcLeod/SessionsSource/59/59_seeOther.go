package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// Initialization ________________________________________________________

var templ *template.Template

func init() {
	templ = template.Must(template.ParseFiles("index-so.gohtml"))
}

// local functions _______________________________________________________

func fo(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Your request method at fo: ", r.Method, "\n \n")
}

func bar(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Your request method at bar: ", r.Method, "\n \n")

	// process form submission here
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

func barred(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Your request method at barred: ", r.Method, "\n \n")
	templ.ExecuteTemplate(w, "index-so.gohtml", http.StatusSeeOther)
}

// Main __________________________________________________________________

func main() {
	http.HandleFunc("/", fo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/barred", barred)

	http.Handle("/favicon.ico", http.NotFoundHandler())

	port := ":8070"

	http.ListenAndServe(port, nil)
}
