package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/satori/uuid"
)

var templat *template.Template

func init() {
	templat = template.Must(template.ParseGlob("templates/*"))
}

func getCookie(w http.ResponseWriter, r *http.Request) *http.Cookie {
	cookie, err := r.Cookie("session")
	if err != nil {
		sessionID := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session",
			Value: sessionID.String(),
		}
		http.SetCookie(w, cookie)
	}
	return cookie
}

func index(w http.ResponseWriter, r *http.Request) {
	cookie := getCookie(w, r)
	templat.ExecuteTemplate(w, "index.gohtml", cookie)
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	port := ":8070"
	fmt.Println("Listening and Serving at port", port)
	http.ListenAndServe(port, nil)
}
