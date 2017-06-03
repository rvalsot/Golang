package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

type cow struct {
	Name string
	Age  int
}

type rabbit struct {
	Name string
	Age  int
}

func initials(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 3 {
		s = s[:3]
	}
	return s
}

var fm = template.FuncMap{
	"upperCase": strings.ToUpper,
	"initials":  initials,
}

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseFiles("tmp.gohtml"))
}

func main() {

	vaquita1 := cow{"Pinta", 1}
	vaquita2 := cow{"Martha", 2}
	vaquita3 := cow{"Petra", 3}

	conejito1 := rabbit{"Conejin", 1}
	conejito2 := rabbit{"Orejon", 3}
	conejito3 := rabbit{"Peludo", 2}

	vaquitas := []cow{vaquita1, vaquita2, vaquita3}
	conejitos := []rabbit{conejito1, conejito2, conejito3}

	dataMap := struct {
		Cows    []cow
		Rabbits []rabbit
	}{
		vaquitas,
		conejitos,
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tmp.gohtml", dataMap)
	if err != nil {
		log.Fatalln(err)
	}

}
