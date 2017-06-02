package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

var templ *template.Template

func init() {
	templ = template.Must(template.ParseFiles("templ.gohtml"))
}

func main() {
	fmt.Println("Class 09 Data into templates")

	data := "mango"
	err := templ.ExecuteTemplate(os.Stdout, "templ.gohtml", data)
	if err != nil {
		log.Fatalln(err)
	}

}
