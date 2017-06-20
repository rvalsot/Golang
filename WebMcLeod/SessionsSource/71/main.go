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

var dbUsers = map[string]user{}      // userID, userData
var dbSessions = map[string]string{} // sessionID, userID

func init() {
	templat = template.Must(template.ParseGlob("templates/*"))

	// Creation of a test user
	ep, _ := bcrypt.GenerateFromPassword([]byte("123"), bcrypt.MinCost)
	dbUsers["test@mail.com"] = user{"test@mail.com", ep, "Vaca", "Pinta"}
}

// Main __________________________________________________________________

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/login", login)
	http.HandleFunc("/singup", singup)

	http.Handle("/favicon.ico", http.NotFoundHandler())

	port := ":8070"
	fmt.Println("Listening and Serving at port ", port)

	http.ListenAndServe(port, nil)
}
