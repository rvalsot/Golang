package main

import (
	"fmt"
)

// ✓ 3.- create the person struct with first & last name ----------------------

type person struct {
	fName string
	lName string
	foods []string
}

// ✓ 6.- Add to person the method walk ----------------------------------------

func (p person) walk() string {

	phrase := p.fName + " " + p.lName + " is walking!"

	return fmt.Sprintln(phrase)
}

// 7.- Create vehicle, sedan, and truck structs -------------------------------

type vehicle struct {
	Color string
	Doors int
}

type truck struct {
	vehicle
	fifthWheel bool
}

type sedan struct {
	vehicle
	luxury bool
}

// ✓ 9.- Transportation device methods that return strings

func (t truck) transportationDevice() string {
	transports := "barrels"
	return transports
}

func (s sedan) transportationDevice() string {
	transports := "people"
	return transports
}

// 10.- Create transportation interface

type transportation interface {
	transportationDevice() string
}

func report(t transportation) string {
	report := fmt.Sprintln(t.transportationDevice())
	return report
}

// func main ------------------------------------------------------------------

func main() {
	fmt.Println("These are exercises of Hands On #2")
	/*
		// ✓ 1.- Create a slice, print it, range printing the index, range printing index & value.

		mySlice := []int{2, 4, 8, 16, 32, 64, 128, 256}
		fmt.Println("The whole slice")
		fmt.Println(mySlice)

		fmt.Println("Printing the indexes")
		for i := range mySlice {
			fmt.Println(i)
		}

		fmt.Println("Printing indexes and values")
		for i, x := range mySlice {
			fmt.Println(i, " : ", x)
		}
	*/

	/*
		// ✓ 2.- Initialize a map key:string & value: int, print it, range over the key, range over the key and value.
		myMap := map[string]int{
			"Finche":   90,
			"Agaporni": 350,
			"Ninfa":    1050,
		}

		fmt.Print("\n")
		fmt.Println("Whole map")
		fmt.Println(myMap)

		fmt.Print("\n")
		fmt.Println("Keys")
		for i := range myMap {
			fmt.Println(i)
		}

		fmt.Print("\n")
		fmt.Println("Keys & Values")
		for i, x := range myMap {
			fmt.Println(i, ": \t $", x)
		}
	*/

	// ✓ 4.- create p1 person, print it
	p1 := person{
		fName: "Niara",
		lName: "the Hedgehog",
		foods: []string{"Worms", "Croquettes", "Crickets", "Bettles"},
	}

	/*
		// ✓ 5.- Add an slice with foods, range-print them.
		fmt.Println(p1.fName, p1.lName, " loves to eat:")
		for _, f := range p1.foods {
			fmt.Println(f)
		}
	*/

	// ✓  6.- Add to person struct walk method, print it out.
	fmt.Println(p1.walk())

	// ✓ 8.- Create a truck, a sedan, print them, print one property.

	camioneta := truck{
		vehicle{"Green",
			2,
		},
		true,
	}

	carro := sedan{
		vehicle{"Blue",
			4,
		},
		true,
	}
	/*
		fmt.Println("Camioneta:", camioneta)
		fmt.Println("Carro:", carro)
		fmt.Println("Car doors", carro.Doors)
		fmt.Println("Does the truck has the fifth wheel?", camioneta.fifthWheel)
	*/

	/*
		// 9.- Transportation devices print
		fmt.Printf("A %T transports: %v \n", camioneta, camioneta.transportationDevice())
		fmt.Printf("A %T transports: %v \n", carro, carro.transportationDevice())
	*/

	// 10.- Interface printing of report function
	fmt.Println(report(carro))
	fmt.Println(report(camioneta))

}
