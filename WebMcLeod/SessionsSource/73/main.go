package main

import (
	"fmt"
	"html/template"
	"net/http"
)

// Definitions ___________________________________________________________

type user struct {
	UserName string
	Password []byte
	Role     string
	First    string
	Last     string
}

var dbUsers = map[string]user{}
var dbSessions = map[string]string{}

// Initalizations ________________________________________________________

var templat *template.Template

func init() {
	templat = template.Must(template.ParseGlob("templates/*"))
}

// Main __________________________________________________________________

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/login", login)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/singup", singup)

	http.Handle("/favicon.ico", http.NotFoundHandler())

	port := ":8070"

	fmt.Println("Listening And Serving at port", port)
	http.ListenAndServe(port, nil)

}
