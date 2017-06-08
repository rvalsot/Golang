package main

import (
	"log"
	"os"
	"text/template"
)

type meal struct {
	Name        string
	Description string
	Class       string
	Cost        float64
}

type menu struct {
	Meals []meal
}

var templat *template.Template

func init() {
	templat = template.Must(template.ParseFiles("template1.gohtml"))
}

func main() {

	dataMap := menu{
		Meals: []meal{
			meal{"Chilaquiles", "Chilaquiles con pollo y chipotle", "Desayuno", 70},
			meal{"Tacos al vapor", "De frijol, papa, carne o requesón", "Desayuno", 10},
			meal{"Molletes", "Dulces o salados", "Desayuno", 20},
			meal{"Jugo de lima", "Jugo de lima natural, 500 mℓ", "Bebida", 15},
			meal{"Refresco", "Refresco, 355 mℓ", "Bebida", 18},
			meal{"Empanadas de huitlacoche", "Empadanadas (3) de huitlacoche", "Comida", 90},
			meal{"Torta", "Con queso, lomo y/o aguacate", "Comida", 35},
			meal{"Churros", "Churros con relleno al gusto (6)", "Cena", 50},
			meal{"Enchiladas", "Con pollo o pato", "Comida", 60},
		},
	}

	err := templat.Execute(os.Stdout, dataMap)
	if err != nil {
		log.Fatalln("Erro while executing:", err)
	}

}
