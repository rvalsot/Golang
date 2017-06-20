package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/satori/uuid"
	"golang.org/x/crypto/bcrypt"
)

// Initialization ________________________________________________________

type user struct {
	UserName string
	Password []byte
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

// func singup(w http.ResponseWriter, r *http.Request) {
// 	if alreadyLoggedIn(r) {
// 		http.Redirect(w, r, "/", http.StatusSeeOther)
// 		return
// 	}
//
// 	// Process form submission
// 	if r.Method == http.MethodPost {
//
// 		// get form the values
// 		un := r.FormValue("username")
// 		p := r.FormValue("password")
// 		f := r.FormValue("firstname")
// 		l := r.FormValue("lastname")
//
// 		// is username already taken?
// 		if _, ok := dbUsers[un]; ok {
// 			http.Error(w, "User name already taken", http.StatusForbidden)
// 			return
// 		}
//
// 		// create session
// 		sessionID := uuid.NewV4()
// 		cookie := &http.Cookie{
// 			Name:  "session",
// 			Value: sessionID.String(),
// 		}
//
// 		http.SetCookie(w, cookie)
// 		dbSessions[cookie.Value] = un
//
// 		// store user in dbUsers
// 		u := user{un, p, f, l}
// 		dbUsers[un] = u
//
// 		// redirect
// 		http.Redirect(w, r, "/", http.StatusSeeOther)
// 		return
// 	}
// 		templat.ExecuteTemplate(w, "singup.gohtml", nil)
//  }

func encryptedSingUp(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	var u user

	// process form submission
	if r.Method == http.MethodPost {
		un := r.FormValue("username")
		p := r.FormValue("password")
		f := r.FormValue("firstname")
		l := r.FormValue("lastname")

		// user taken?

		if _, ok := dbUsers[un]; ok {
			http.Error(w, "Username already taken", http.StatusForbidden)
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
		ep, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		u = user{un, ep, f, l}
		dbUsers[un] = u

		//redirect
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	templat.ExecuteTemplate(w, "singup.gohtml", u)
}

// Main section __________________________________________________________

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/singup", encryptedSingUp)

	port := ":8070"
	fmt.Println("Listening And Serving at port ", port)

	http.ListenAndServe(port, nil)
}
