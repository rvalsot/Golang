package main

import (
	"net/http"

	"github.com/satori/uuid"
	"golang.org/x/crypto/bcrypt"
)

// Index handler _________________________________________________________

func index(w http.ResponseWriter, r *http.Request) {
	u := getUser(w, r)
	templat.ExecuteTemplate(w, "index.gohtml", u)
}

// Bar handler ___________________________________________________________

func bar(w http.ResponseWriter, r *http.Request) {
	u := getUser(w, r)

	if !alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	if u.Role != "007" {
		restriction := "You must be in special forces to join this bar"
		http.Error(w, restriction, http.StatusForbidden)
		return
	}
	templat.ExecuteTemplate(w, "bar.gohtml", u)
}

// Login handler _________________________________________________________

func login(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	var u user

	// process form submission
	if r.Method == http.MethodPost {
		un := r.FormValue("user_name")
		ps := r.FormValue("password")

		// is there the user name?
		u, ok := dbUsers[un]
		if !ok {
			announce := "UserName or password incorrect"
			http.Error(w, announce, http.StatusForbidden)
			return
		}

		// does the password matches?
		err := bcrypt.CompareHashAndPassword(u.Password, []byte(ps))
		if err != nil {
			announcement := "Password do not match"
			http.Error(w, announcement, http.StatusForbidden)
			return
		}

		// create session cookie
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
	templat.ExecuteTemplate(w, "login.gohtml", u)

}

// Logout handler ________________________________________________________

func logout(w http.ResponseWriter, r *http.Request) {
	if !alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// get the cookie
	cookie, _ := r.Cookie("session")

	// delete the session
	delete(dbSessions, cookie.Value)

	// remove the cookie
	cookie = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie) // or unset

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

// Singup handler ________________________________________________________

func singup(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	var u user

	// process singup form submission
	if r.Method == http.MethodPost {
		un := r.FormValue("user_name")
		ps := r.FormValue("password")
		ro := r.FormValue("role")
		fi := r.FormValue("first_name")
		la := r.FormValue("last_name")

		// user_name already taken?
		if _, ok := dbUsers[un]; ok {
			errMess := "user name already taken"
			http.Error(w, errMess, http.StatusForbidden)
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
		ep, err := bcrypt.GenerateFromPassword([]byte(ps), bcrypt.MinCost)
		if err != nil {
			errMess := "Internal server error"
			http.Error(w, errMess, http.StatusInternalServerError)
			return
		}
		u = user{un, ep, ro, fi, la}
		dbUsers[un] = u

		// redirect
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	templat.ExecuteTemplate(w, "singup.gohtml", u)
}
