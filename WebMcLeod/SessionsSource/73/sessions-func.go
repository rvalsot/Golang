package main

import (
	"net/http"

	"github.com/satori/uuid"
)

func getUser(w http.ResponseWriter, r *http.Request) user {
	// get the cookie

	cookie, err := r.Cookie("session")
	if err != nil {
		sessionID := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session",
			Value: sessionID.String(),
		}
	}
	http.SetCookie(w, cookie)

	// if the user exists al ready, get the user
	var usr user
	if un, ok := dbSessions[cookie.Value]; ok {
		usr = dbUsers[un]
	}
	return usr
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
