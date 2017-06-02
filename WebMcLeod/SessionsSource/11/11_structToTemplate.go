package main

import (
	"log"
	"os"
	"text/template"
)

var temp *template.Template

// Candy is a struct that has a Name string and "yes/no" string to tell us if it has Chocolate
type Candy struct {
	Name      string
	Chocolate string
}

func init() {
	temp = template.Must(template.ParseFiles("template3.html"))
}

func main() {

	dataMap := Candy{
		"doughnut",
		"yes",
	}

	err := temp.Execute(os.Stdout, dataMap)
	if err != nil {
		log.Fatalln(err)
	}
}
