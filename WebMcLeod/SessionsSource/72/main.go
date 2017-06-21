package main

import (
	"fmt"
	"html/template"
	"net/http"

	"golang.org/x/crypto/bcrypt"
)

// Definitions ___________________________________________________________

type user struct {
	UserName string
	Password []byte
	First    string
	Last     string
}

// Initialization ________________________________________________________

var templat *template.Template
var dbUsers = map[string]user{}
var dbSessions = map[string]string{}

func init() {
	templat = template.Must(template.ParseGlob("templates/*"))
	ep, _ := bcrypt.GenerateFromPassword([]byte("123"), bcrypt.MinCost)
	dbUsers["vaquita@3.com"] = user{"vaquita@3.com", ep, "Vaquita", "3s"}
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
