package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func main() {
	fmt.Println("File 08 Parsing HTML Go Source into a file")

	filename := "08_base.gohtml"

	// Get the file contents into the template
	templ, err := template.ParseFiles(filename)
	if err != nil {
		log.Fatalln("Ops! → ", err)
	}

	newFile, err := os.Create("08_index.html")
	if err != nil {
		log.Println("Error creating file", err)
	}
	defer newFile.Close()

	err = templ.Execute(newFile, nil)
	if err != nil {
		log.Fatalln(err)
	}

}
