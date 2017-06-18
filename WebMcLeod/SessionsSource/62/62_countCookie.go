package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", fus)
	http.Handle("/favicon.ico", http.NotFoundHandler())

	port := ":8070"
	fmt.Println("Listened and Served at port", port)
	http.ListenAndServe(port, nil)
}

func fus(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("niara-cookie")
	if err != http.ErrNoCookie {
		cookie = &http.Cookie{
			Name:  "niara-cookie",
			Value: "0",
		}
	}

	count, err := strconv.Atoi(cookie.Value)
	if err != nil {
		log.Fatalln(err)
	}

	count++
	cookie.Value = strconv.Itoa(count)
	http.SetCookie(w, cookie)

	io.WriteString(w, cookie.Value)
	fmt.Fprint(w, "\n Counter: ", cookie.Value)
}
