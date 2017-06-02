package main

import (
	"log"
	"os"
	"text/template"
)

var temp *template.Template

func init() {
	temp = template.Must(template.ParseFiles("template2.html"))
}

func main() {

	dataMap := map[string]string{
		"Cheesecake":  "$100",
		"Moka":        "$50",
		"Carrot cake": "$75",
	}

	err := temp.Execute(os.Stdout, dataMap)
	if err != nil {
		log.Fatalln(err)
	}
}
