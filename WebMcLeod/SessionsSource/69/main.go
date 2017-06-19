package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/satori/uuid"
)

// Initialization ________________________________________________________

type user struct {
	UserName string
	Password string
	First    string
	Last     string
}

var templat *template.Template

var dbUsers = map[string]user{}      // map of userID & user struct
var dbSessions = map[string]string{} // map of sessionID & userID

func init() {
	templat = template.Must(template.ParseGlob("templates/*"))
}

// Local functions _______________________________________________________

func index(w http.ResponseWriter, r *http.Request) {
	u := getUser(w, r)
	templat.ExecuteTemplate(w, "index.gohtml", u)
}

func bar(w http.ResponseWriter, r *http.Request) {
	u := getUser(w, r)
	if !alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	templat.ExecuteTemplate(w, "bar.gohtml", u)
}

func singup(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Process form submission
	if r.Method == http.MethodPost {

		// get form the values
		un := r.FormValue("username")
		p := r.FormValue("password")
		f := r.FormValue("firstname")
		l := r.FormValue("lastname")

		// is username already taken?
		if _, ok := dbUsers[un]; ok {
			http.Error(w, "User name already taken", http.StatusForbidden)
			return
		}

		// create session
		sessionID := uuid.NewV4()
		cookie := &http.Cookie{
			Name:  "session",
			Value: sessionID.String(),
		}

		http.SetCookie(w, cookie)
		dbSessions[cookie.Value] = un

		// store user in dbUsers
		u := user{un, p, f, l}
		dbUsers[un] = u

		// redirect
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	templat.ExecuteTemplate(w, "singup.gohtml", nil)
}

// Main section __________________________________________________________

func main() {

	port := ":8070"
	fmt.Println("Listening And Serving at port ", port)

	http.ListenAndServe(port, nil)
}
