package main

import (
	"log"
	"net/http"
	"path/filepath"
	"sync"
	"text/template"
)

type templateHandler struct {
	once     sync.Once
	filename string
	templ    *template.Template
}

func (t *templateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	t.once.Do(func() {
		t.templ = template.Must(
			template.ParseFiles(
				filepath.Join("templates", t.filename)))
	})
	t.templ.Execute(w, nil)
}

func main() {

	myRoom := newRoom()

	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", myRoom)

	// get the room going
	go myRoom.run()

	// start the web server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe broken:", err)
	}

	//
	//
	// http.Handle("/", &templateHandler{filename: "farm.html"})
	//
	// if err := http.ListenAndServe(":3000", nil); err != nil {
	// 	log.Fatal("ListenAndServe Fail:", err)
	// }
}
