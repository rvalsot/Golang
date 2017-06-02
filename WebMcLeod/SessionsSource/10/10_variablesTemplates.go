package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var templ *template.Template

func init() {
	templ = template.Must(template.ParseFiles("template.gohtml"))
}

func main() {
	fmt.Println("Class 10: Printing variables into templates")

	file := "template.gohtml"
	data := "will be transformed into ketchup."
	err := templ.ExecuteTemplate(os.Stdout, file, data)
	if err != nil {
		log.Fatalln(err)
	}
}
