package main

import (
	"net/http"

	"github.com/satori/uuid"
)

// Get the user, via its cookie

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

	// get the user if already exists
	var usr user
	if un, ok := dbSessions[cookie.Value]; ok {
		usr = dbUsers[un]
	}

	return usr

}

// Already logged in?

func alreadyLoggedIn(r *http.Request) bool {
	// this will return an error if the cookie we're looking for does not exists
	cookie, err := r.Cookie("session")
	if err != nil {
		return false
	}

	un := dbSessions[cookie.Value]
	_, ok := dbUsers[un]
	return ok

}
