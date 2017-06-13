package main

import (
	"html/template"
	"log"
	"os"
)

var templ *template.Template

// Page is a struct to add elements to a page
type Page struct {
	Title   string
	Heading string
	Input   string
}

func init() {
	templ = template.Must(template.ParseFiles("template2.gohtml"))
}

func main() {

	home := Page{
		Title:   "Escaped",
		Heading: "Danger is scaped with html/template, not with text/template",
		Input:   `<script> alert("Ya get catched!")</script>`,
	}

	// Note that here we're using html and not text /template, so we expect our injection to be stopped.
	err := templ.ExecuteTemplate(os.Stdout, "template2.gohtml", home)
	if err != nil {
		log.Fatalln(err)
	}
}
