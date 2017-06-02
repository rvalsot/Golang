package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var temp *template.Template

func init() {
	temp = template.Must(template.ParseFiles("template.html"))
}

func main() {
	fmt.Println("11: slice of strings into template")

	animals := []string{"cow", "pig", "chicken", "rabit", "goat"}
	err := temp.Execute(os.Stdout, animals)
	if err != nil {
		log.Fatalln(err)
	}
}
