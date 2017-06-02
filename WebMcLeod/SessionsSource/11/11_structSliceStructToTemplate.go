package main

import (
	"log"
	"os"
	"text/template"
)

var temp *template.Template

// House is a struct that describes a house
type House struct {
	Address string
	Area    int
	Rooms   int
}

// ApartmentBuilding is a strut that describes an apartment building
type ApartmentBuilding struct {
	Address          string
	ApartmentsNumber int
	Levels           int
}

// Properties is a struct that lists possessions
type Properties struct {
	Houses             []House
	ApartmentBuildings []ApartmentBuilding
}

func init() {
	temp = template.Must(template.ParseFiles("template5.html"))
}

func main() {

	SJL := House{"SJL", 300, 6}
	GDL := House{"GDL", 75, 3}

	LEO := ApartmentBuilding{"LEO", 8, 4}
	PVA := ApartmentBuilding{"PVA", 6, 3}

	houses := []House{SJL, GDL}
	apartmentBuildings := []ApartmentBuilding{LEO, PVA}

	dataMap := Properties{
		Houses:             houses,
		ApartmentBuildings: apartmentBuildings,
	}

	err := temp.Execute(os.Stdout, dataMap)
	if err != nil {
		log.Fatalln(err)
	}

	/*
		newFilename := "toIndex.html"
		newFile, err := os.Create(newFilename)
		if err != nil {
			log.Println("File could not be created: \n", err)
		}
		defer newFile.Close()

		err = temp.Execute(newFile, nil)
		if err != nil {
			log.Fatalln(err)
		}.
	*/

}
