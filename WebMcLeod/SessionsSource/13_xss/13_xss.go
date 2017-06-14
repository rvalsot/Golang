package main

import (
	"log"
	"os"
	"text/template"
)

// Page is a basic page
type Page struct {
	Title   string
	Heading string
	Input   string
}

var templat *template.Template

func init() {
	templat = template.Must(template.ParseFiles("template1.gohtml"))
}

func main() {
	home := Page{
		Title:   "Nothing Escaped",
		Heading: "Nothing Escaped with text/template",
		Input:   `<script>alert("Yo!")</script>`,
	}

	err := templat.ExecuteTemplate(os.Stdout, "template1.gohtml", home)
	if err != nil {
		log.Fatalln(err)
	}
}
