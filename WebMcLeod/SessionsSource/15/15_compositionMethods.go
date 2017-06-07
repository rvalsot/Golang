package main

import (
	"log"
	"os"
	"text/template"
)

type course struct {
	Number, Name, Units string
}

type quarter struct {
	Term    string
	Courses []course
}

type year struct {
	Primavera, Verano, Otono quarter
}

var templ *template.Template

func init() {
	templ = template.Must(template.ParseFiles("template2.gohtml"))
}

func main() {
	myYear := year{
		Primavera: quarter{
			Term: "Primavera",
			Courses: []course{
				course{"01", "Go Intro", "8"},
				course{"02", "Go Web", "8"},
				course{"03", "Go Mobile", "8"},
			},
		},
		Verano: quarter{
			Term: "Verano",
			Courses: []course{
				course{"04", "Go Database", "8"},
			},
		},
		Otono: quarter{
			Term: "Otono",
			Courses: []course{
				course{"05", "Go Connections", "8"},
				course{"06", "Go DevOps", "8"},
				course{"07", "Go Production", "8"},
			},
		},
	}

	err := templ.Execute(os.Stdout, myYear)
	if err != nil {
		log.Fatalln(err)
	}

}
