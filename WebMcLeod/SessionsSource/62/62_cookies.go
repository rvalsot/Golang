package main

import (
	"fmt"
	"net/http"
)

func set(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "niara-cookie",
		Value: "wormy cookie",
	})

	fmt.Fprintln(w, "Cookie Written, chek browser")
	fmt.Fprintln(w, "Dev/tools/application/cookies")
}

func read(w http.ResponseWriter, r *http.Request) {
	coo, err := r.Cookie("niara-cookie")
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	fmt.Fprintln(w, "Your cookie: ", coo)
}

func main() {
	http.HandleFunc("/", set)
	http.HandleFunc("/read", read)

	http.Handle("/favicon.ico", http.NotFoundHandler())

	port := ":8070"
	fmt.Println("Listening and Serving at port ", port)
	http.ListenAndServe(port, nil)
}
