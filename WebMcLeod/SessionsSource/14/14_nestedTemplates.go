package main

import (
	"log"
	"os"
	"text/template"
)

var templat *template.Template

func init() {
	templat = template.Must(template.ParseGlob("./*.gohtml"))
}

func main() {
	data := "is just covfefe"
	err := templat.ExecuteTemplate(os.Stdout, "index1.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}
}
