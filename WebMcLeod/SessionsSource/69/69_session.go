package main

import (
	"net/http"

	"github.com/satori/uuid"
)

func getUser(w http.ResponseWriter, r *http.Request) user {
	// get cookie
	cookie, err := r.Cookie("session")
	if err != nil {
		sessionID := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session",
			Value: sessionID.String(),
		}
	}

	http.SetCookie(w, cookie)

	// if the user, already exists, get it

	var u user
	if un, ok := dbSessions[cookie.Value]; ok {
		u = dbUsers[un]
	}
	return u
}

func alreadyLoggedIn(r *http.Request) bool {
	cookie, err := r.Cookie("session")
	if err != nil {
		return false
	}
	un := dbSessions[cookie.Value]
	_, ok := dbUsers[un]
	return ok
}
