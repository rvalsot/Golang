package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func main() {
	fmt.Println("File 08 Parsing HTML")

	filename := "08_base.gohtml"

	// Get the file contents into the template
	templ, err := template.ParseFiles(filename)
	if err != nil {
		log.Fatalln("Ops! → ", err)
	}

	err = templ.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln("Error: \n", err)
	}

}
