package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

func main() {
	fmt.Println("File 08 Parsing Files Sources into a CLI")

	templ, err := template.ParseGlob("templates/*")
	if err != nil {
		log.Fatalln("Error found! ", err)
	}

	err = templ.Execute(os.Stdout, nil)
	if err != nil {
		log.Fatalln("Execute error: ", err)
	}

	fmt.Println("Individual execution:")

	err = templ.ExecuteTemplate(os.Stdout, "duke250.gohtml", nil)
	if err != nil {
		log.Fatalln("Error with Duke 250: \n", err)
	}

	err = templ.ExecuteTemplate(os.Stdout, "himalayan.gohtml", nil)
	if err != nil {
		log.Fatalln("Error with Himalayan: \n", err)
	}

	err = templ.ExecuteTemplate(os.Stdout, "speedmaster.gohtml", nil)
	if err != nil {
		log.Fatalln("Error with Speedmaster: \n", err)
	}

}
