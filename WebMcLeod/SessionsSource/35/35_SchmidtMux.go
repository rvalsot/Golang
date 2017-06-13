package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

var templat *template.Template

func init() {
	templat = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	mux := httprouter.New()

	mux.GET("/", index)
	mux.GET("/about", about)
	mux.GET("/contact", contact)
	mux.GET("/apply", apply)
	mux.POST("/apply", applyProcess)
	mux.GET("/user/:name", user)
	mux.GET("/blog/:category/:article", blogRead)
	mux.POST("/blog/:category/:article", blogWrite)

	http.ListenAndServe(":8079", mux)
}

func user(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "USER %s \n", p.ByName("name"))
}

func blogRead(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "READ CATEGORY %s \n", p.ByName("category"))
	fmt.Fprintf(w, "READ ARTICLE %s \n", p.ByName("article"))
}

func blogWrite(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	fmt.Fprintf(w, "WRITE CATEGORY, %s \n", p.ByName("category"))
	fmt.Fprintf(w, "WRITE ARTICLE, %s \n", p.ByName("article"))
}

func index(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := templat.ExecuteTemplate(w, "index.gohtml", nil)
	HandleError(w, err)
}

func about(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := templat.ExecuteTemplate(w, "about.gohtml", nil)
	HandleError(w, err)
}

func contact(w http.ResponseWriter, req *http.Request, _ httprouter.Params) {
	err := templat.ExecuteTemplate(w, "contact.gohtml", nil)
	HandleError(w, err)
}

func apply(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := templat.ExecuteTemplate(w, "apply.gohtml", nil)
	HandleError(w, err)
}

func applyProcess(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	err := templat.ExecuteTemplate(w, "applyProcess.gohtml", nil)
	HandleError(w, err)
}

// HandleError helps us whenever there's a problem with our web petitions
func HandleError(w http.ResponseWriter, err error) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Fatalln(err)
	}
}
