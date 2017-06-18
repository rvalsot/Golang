package main

import (
	"fmt"
	"net/http"
)

// Initialization ________________________________________________________

func fooo(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Your request at fooo: ", r.Method, "\n \n")
}

func barrr(w http.ResponseWriter, r *http.Request) {
	fmt.Print("Your request at barrr: ", r.Method, "\n \n")
	http.Redirect(w, r, "/", http.StatusMovedPermanently)
}

// Main __________________________________________________________________

func main() {
	http.HandleFunc("/", fooo)
	http.HandleFunc("/barrr", barrr)

	http.Handle("/favicon.ico", http.NotFoundHandler())

	port := ":8070"
	fmt.Println("Listening and Serving port", port)
	http.ListenAndServe(port, nil)
}
