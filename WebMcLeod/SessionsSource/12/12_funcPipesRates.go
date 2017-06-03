package main

import (
	"log"
	"math"
	"os"
	"text/template"
)

var templ *template.Template

func init() {
	templ = template.Must(template.New("").Funcs(funcMap).ParseFiles("template3.gohtml"))
}

var funcMap = template.FuncMap{
	"second": second,
	"third":  third,
	"fourth": fourth,
}

func second(x float64) float64 {
	return x * x
}

func third(x float64) float64 {
	return math.Pow(x, 3)
}

func fourth(x float64) float64 {
	return math.Pow(x, 4)
}

func main() {
	err := templ.ExecuteTemplate(os.Stdout, "template3.gohtml", 3.2)
	if err != nil {
		log.Fatalln(err)
	}
}
