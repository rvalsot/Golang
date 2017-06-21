package main

import (
	"net/http"

	"github.com/satori/uuid"
	"golang.org/x/crypto/bcrypt"
)

// Index handler _________________________________________________________

func index(w http.ResponseWriter, r *http.Request) {
	usr := getUser(w, r)
	templat.ExecuteTemplate(w, "index.gohtml", usr)
}

// Bar handler ___________________________________________________________

func bar(w http.ResponseWriter, r *http.Request) {
	usr := getUser(w, r)
	if !alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	templat.ExecuteTemplate(w, "bar.gohtml", usr)
}

// Login handler _________________________________________________________

func login(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	var usr user

	// process login form submission
	if r.Method == http.MethodPost {
		un := r.FormValue("user_name")
		ps := r.FormValue("password")

		// user_name exists?
		usr, ok := dbUsers[un]
		if !ok {
			http.Error(w, "Username and/or password do not match", http.StatusForbidden)
			return
		}

		// entered password and stored one match?
		err := bcrypt.CompareHashAndPassword(usr.Password, []byte(ps))
		if err != nil {
			http.Error(w, "Password do not match", http.StatusForbidden)
			return
		}

		// session creation
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

	templat.ExecuteTemplate(w, "login.gohtml", usr)
}

// Logout handler ________________________________________________________

func logout(w http.ResponseWriter, r *http.Request) {
	if !alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	cookie, _ := r.Cookie("session")

	// delete the session associated with the cookie
	delete(dbSessions, cookie.Value)

	// remove the cookie with a timeout
	cookie = &http.Cookie{
		Name:   "session",
		Value:  "",
		MaxAge: -1,
	}
	http.SetCookie(w, cookie)

	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

// Sing up handler _______________________________________________________

func singup(w http.ResponseWriter, r *http.Request) {
	if alreadyLoggedIn(r) {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	var usr user

	// process singup form submission
	if r.Method == http.MethodPost {
		// get form values
		un := r.FormValue("user_name")
		ps := r.FormValue("password")
		fi := r.FormValue("first_name")
		la := r.FormValue("last_name")

		// username taken?
		if _, ok := dbUsers[un]; ok {
			http.Error(w, "Username already taken", http.StatusForbidden)
		}

		// create user's session
		sessionID := uuid.NewV4()
		cookie := &http.Cookie{
			Name:  "session",
			Value: sessionID.String(),
		}
		http.SetCookie(w, cookie)

		dbSessions[cookie.Value] = un

		// store the new user in dbUsers
		ep, err := bcrypt.GenerateFromPassword([]byte(ps), bcrypt.MinCost)
		if err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}

		usr = user{un, ep, fi, la}
		dbUsers[un] = usr
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	templat.ExecuteTemplate(w, "singup.gohtml", usr)
}
