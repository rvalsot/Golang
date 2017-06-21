package main

import (
	"fmt"
	"html/template"
	"net/http"
	"strings"

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

func appendValue(w http.ResponseWriter, c *http.Cookie) *http.Cookie {
	// values
	p1 := "niara1.jpeg"
	p2 := "niara2.jpeg"
	p3 := "niara3.jpeg"

	// append
	s := c.Value
	if !strings.Contains(s, p1) {
		s += "|" + p1
	}
	if !strings.Contains(s, p2) {
		s += "|" + p2
	}
	if !strings.Contains(s, p3) {
		s += "|" + p3
	}
	c.Value = s
	http.SetCookie(w, c)
	return c
}

func index(w http.ResponseWriter, r *http.Request) {
	c := getCookie(w, r)
	c = appendValue(w, c)
	xs := strings.Split(c.Value, "|")

	templat.ExecuteTemplate(w, "index.gohtml", xs)
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	port := ":8070"
	fmt.Println("Listening and Serving at port", port)
	http.ListenAndServe(port, nil)
}
