package main

/* This file contains the handler functions:
 *  • index
 *  • bar
 *  • login
 *  • singup
 */

import (
	"net/http"

	"github.com/satori/uuid"
	"golang.org/x/crypto/bcrypt"
)

// Index page handler ____________________________________________________

func index(w http.ResponseWriter, r *http.Request) {
	usr := getUser(w, r)
	templat.ExecuteTemplate(w, "index.gohtml", usr)
}

// Bar page handler ______________________________________________________

func bar(w http.ResponseWriter, r *http.Request) {

	usr := getUser(w, r)

	if !alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	templat.ExecuteTemplate(w, "bar.gohtml", usr)
}

// Login page handler ____________________________________________________

func login(w http.ResponseWriter, r *http.Request) {

	// check for loggin
	if alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
	}

	// process form submission
	if r.Method == http.MethodPost {
		un := r.FormValue("user_name")
		ps := r.FormValue("password")

		// is there a user name?
		usr, ok := dbUsers[un]
		if !ok {
			http.Error(w, "User name and/or password do not match", http.StatusForbidden)
			return
		}

		// is password correct?
		err := bcrypt.CompareHashAndPassword(usr.Password, []byte(ps))
		if err != nil {
			http.Error(w, "User name and/or password do not match", http.StatusForbidden)
			return
		}

		// create a session
		sessionID := uuid.NewV4()
		cookie := &http.Cookie{
			Name:  "session",
			Value: sessionID.String(),
		}

		http.SetCookie(w, cookie)
		dbSessions[cookie.Value] = un
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return

	}

	templat.ExecuteTemplate(w, "login.gohtml", nil)

}

// Sing up page __________________________________________________________

func singup(w http.ResponseWriter, r *http.Request) {

	// check for loggin
	if alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	var usr user

	//process form submission
	if r.Method == http.MethodPost {
		// get form values, create user
		un := r.FormValue("user_name")
		ps := r.FormValue("password")
		fi := r.FormValue("first_name")
		la := r.FormValue("last_name")

		// username taken?
		if _, ok := dbUsers[un]; ok {
			http.Error(w, "User name already taken", http.StatusForbidden)
			return
		}

		// create a session
		sessionID := uuid.NewV4()
		cookie := &http.Cookie{
			Name:  "session",
			Value: sessionID.String(),
		}

		http.SetCookie(w, cookie)
		dbSessions[cookie.Value] = un

		// store the user
		ep, err := bcrypt.GenerateFromPassword([]byte(ps), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		usr = user{un, ep, fi, la}
		dbUsers[un] = usr

		// redirect to index
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	templat.ExecuteTemplate(w, "singup.gohtml", usr)

}
