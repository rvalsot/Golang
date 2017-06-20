package main

import (
	"fmt"
	"net/http"

	"github.com/satori/uuid"
)

func index(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		id := uuid.NewV4()
		cookie = &http.Cookie{
			Name:  "session",
			Value: id.String(),
			// Secure: true,
			HttpOnly: true,
			Path:     "/",
		}
		http.SetCookie(w, cookie)
	}

	fmt.Println(cookie)
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("favicon.ico", http.NotFoundHandler())

	port := ":8070"
	fmt.Println("Listening and Serving at port", port)
	http.ListenAndServe(port, nil)

}
