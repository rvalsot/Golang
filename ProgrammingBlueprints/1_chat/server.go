package main

import (
	"flag"
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

// ____________________________________________________________________________

func main() {

	var address = flag.String("address", ":8080", "The address of our web application.")
	flag.Parse() // Parsing of the flag

	myRoom := newRoom()

	http.Handle("/", &templateHandler{filename: "chat.html"})
	http.Handle("/room", myRoom)

	// get the room going
	go myRoom.run()

	// start the web server
	log.Println("Starting web service at: ", *address)
	if err := http.ListenAndServe(*address, nil); err != nil {
		log.Fatal("ListenAndServe broken:", err)
	}

}
