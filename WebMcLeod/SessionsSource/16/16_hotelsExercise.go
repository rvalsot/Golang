package main

import (
	"fmt"
	"log"
	"os"
	"text/template"
)

type hotel struct {
	Name, Address, City, Zip string
	Rooms                    int
	Price                    float64
}

var templat *template.Template

func init() {
	templat = template.Must(template.ParseFiles("template1.gohtml"))
}

func main() {

	dataMap := []hotel{
		hotel{"A", "Vaca 3", "SJL", "395", 25, 700},
		hotel{"B", "Vaca 10", "SJL", "395", 18, 900},
		hotel{"C", "Zarzamora 7", "GDL", "331", 40, 700},
		hotel{"D", "Arandano 2", "GDL", "331", 60, 900},
		hotel{"E", "Calabaza 15", "GDL", "331", 120, 1200},
		hotel{"F", "Maize 10", "ZAM", "351", 10, 700},
	}

	fmt.Println(dataMap)

	err := templat.Execute(os.Stdout, dataMap)
	if err != nil {
		log.Fatalln("Error:", err)
	}

}
