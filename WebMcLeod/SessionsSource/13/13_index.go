package main

import (
	"log"
	"os"
	"text/template"
)

var tmplate *template.Template

func init() {
	tmplate = template.Must(template.ParseFiles("t1.gohtml"))
}

func main() {
	xs := []int{1, 2, 3, 4, 5}

	err := tmplate.Execute(os.Stdout, xs)
	if err != nil {
		log.Fatalln(err)
	}
}
