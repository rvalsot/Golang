package main

import (
	"log"
	"os"
	"text/template"
	"time"
)

var templ *template.Template

func init() {
	templ = template.Must(template.New("").Funcs(funcMap).ParseFiles("template2.gohtml"))
}

var funcMap = template.FuncMap{
	"americanDate": americanDate,
}

func americanDate(t time.Time) string {
	timeStringified := t.Format("Monday, Jan 02, 2006")
	return timeStringified
}

func main() {
	err := templ.ExecuteTemplate(os.Stdout, "template2.gohtml", time.Now())
	if err != nil {
		log.Fatalln(err)
	}
}
