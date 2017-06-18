package main

import (
	"fmt"
	"log"
	"net/http"
)

// Local functions ________________________________________________________

func setter(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "niara-cookie",
		Value: "yuum, wormy cookie",
	})

	fmt.Fprintln(w, "Your cookie is written, check at browser")
	fmt.Fprintln(w, "At Firefox: options/privacy: cookies")
}

func readerr(w http.ResponseWriter, r *http.Request) {
	co1, err := r.Cookie("niara-cookie")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "Your cookie #1: ", co1)
	}

	co2, err := r.Cookie("general")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "Your cookie #2: ", co2)
	}

	co3, err := r.Cookie("specific")
	if err != nil {
		log.Println(err)
	} else {
		fmt.Fprintln(w, "Your cookie #3: ", co3)
	}

}

func abundance(w http.ResponseWriter, r *http.Request) {
	http.SetCookie(w, &http.Cookie{
		Name:  "general",
		Value: "general cookies co.",
	})

	http.SetCookie(w, &http.Cookie{
		Name:  "specific",
		Value: "specific cookies co.",
	})

	fmt.Fprintln(w, "Abundant cookies written, check browser")
	fmt.Fprintln(w, "Firefox: preferences/privacy: cookies")
}

// Main __________________________________________________________________

func main() {
	http.HandleFunc("/", setter)
	http.HandleFunc("/read", readerr)
	http.HandleFunc("/abundance", abundance)

	http.Handle("/favicon.ico", http.NotFoundHandler())

	port := ":8070"
	fmt.Println("Listening and Serving at port ", port)
	http.ListenAndServe(port, nil)
}
