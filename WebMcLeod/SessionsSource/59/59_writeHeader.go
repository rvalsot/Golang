package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var templat *template.Template

func init() {
	templat = template.Must(template.ParseFiles("index-wh.gohtml"))
}

func fu(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Your request method at fu: ", r.Method, "\n \n")
	w.Header().Set("Location", "/")
	w.WriteHeader(http.StatusSeeOther)
}

func bared(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Your request method at bared: ", r.Method, "\n \n")
	templat.ExecuteTemplate(w, "index-wh.gohtml", nil)
}

// Main __________________________________________________________________

func main() {
	http.HandleFunc("/", fu)
	http.HandleFunc("/bared", bared)

	http.Handle("/favicon.ico", http.NotFoundHandler())

	port := ":8070"
	fmt.Println("Listening and Serving at port", port)

	http.ListenAndServe(port, nil)
}
