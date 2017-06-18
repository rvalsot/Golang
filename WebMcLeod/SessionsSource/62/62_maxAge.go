package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", index)
	http.HandleFunc("/sets", sets)
	http.HandleFunc("/reads", reads)
	http.HandleFunc("/expires", expires)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	port := ":8070"
	fmt.Println("Listening and Serving at port ", port)
	http.ListenAndServe(port, nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, `
    <h1> <a href="/sets"> Set a cookie!</a> </h1>
  `)
}

func sets(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "session",
		Value: "some value",
	})
	fmt.Fprintln(w, `
    <h1> <a href="/reads"> Read! </a> </h1>
  `)
}

func reads(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		http.Redirect(w, r, "/set", http.StatusSeeOther)
		return
	}

	fmt.Fprintf(w, `
      <h1>Your cookie:
      <br />
        %v
      </h1>
      <h1>
        <a href="/expire"> Expire </a>
      </h1>
    `, cookie)
}

func expires(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		http.Redirect(w, r, "/set", http.StatusSeeOther)
		return
	}

	// delete our cookie
	cookie.MaxAge = -1
	http.SetCookie(w, cookie)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
