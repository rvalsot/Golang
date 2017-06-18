package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// Initialization ________________________________________________________

var templa *template.Template

func init() {
	templa = template.Must(template.ParseFiles("index-tr.gohtml"))
}

// Local functions _______________________________________________________

func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Your request at method foo: ", r.Method, "\n \n")
}

func barr(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Your request method at barr: ", r.Method, "\n \n")
	http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
}

func barredd(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Your request method at barred: ", r.Method, "\n \n")
	templa.ExecuteTemplate(w, "index-tr.gohtml", nil)
}

// Main _________________________________________________________________

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/barr", barr)
	http.HandleFunc("/barredd", barredd)

	http.Handle("/favicon.ico", http.NotFoundHandler())

	port := ":8070"

	fmt.Println("Listening and Serving at port", port)
	http.ListenAndServe(port, nil)
}
