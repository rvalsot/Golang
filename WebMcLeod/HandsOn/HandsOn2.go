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

// 12.- Type gator
type gator int

// 15.- Method for type gator
func (g gator) greeting() string {
	returned := "Hi, I'm a gator"
	return returned
}

// 16.- Type flamingo, interfaces and methods associated

type flamingo int

func (f flamingo) greeting() string {
	returned := "Hi, I'm a flamingo"
	return returned
}

type swampGuys interface {
	greeting() string
}

func bayou(sg swampGuys) {
	fmt.Println(sg.greeting())
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

	/*
		// ✓ 4.- create p1 person, print it
		p1 := person{
			fName: "Niara",
			lName: "the Hedgehog",
			foods: []string{"Worms", "Croquettes", "Crickets", "Bettles"},
		}
	*/

	/*
		// ✓ 5.- Add an slice with foods, range-print them.
		fmt.Println(p1.fName, p1.lName, " loves to eat:")
		for _, f := range p1.foods {
			fmt.Println(f)
		}
	*/

	/*
		// ✓  6.- Add to person struct walk method, print it out.
		fmt.Println(p1.walk())
	*/

	/*
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

	/*
		// 10.- Interface printing of report function
		fmt.Println(report(carro))
		fmt.Println(report(camioneta))
	*/

	// 11.- Creation of type gator

	var g1 gator = 1
	fmt.Println(g1)
	fmt.Printf("%T \n", g1)

	// 12, 13, 14.- Creation of an int, g1 "gator" value assignment to it, did it broke? Conversion

	var x = int(g1)
	fmt.Println(x)
	fmt.Printf("%T \n", x)

	// 15.- Method for type gator calling
	fmt.Println(g1.greeting())

	var f1 flamingo

	// 16.- Application of the interface
	bayou(f1)
	bayou(g1)

	// 17.- Bit stuff

	s := "Sorry but flamingos prefer to eat shrimp."

	// 17.1.- Print s
	fmt.Println(s)

	// 17.2.- Print s into slice of byte

	sSlice := []byte(s)
	fmt.Println(sSlice)

	// 17.3.- Print to slice of byte and to string again

	s2 := string(sSlice)
	fmt.Println(s2)

	// 17.4.- Slice "Sorry but"
	le := len("Sorry but")
	for k, v := range sSlice {
		if k <= le {
			fmt.Print(string(v))
		}
	}
	fmt.Println("")

	// 17.5.- Slice "flamingos prefer to"
	le2 := len("flamingos prefer to")
	for k, v := range sSlice {
		if k >= le && k <= le+le2 {
			fmt.Print(string(v))
		}
	}
	fmt.Println("")

	// 17.6.- Slice "eat shrimp."
	le3 := len("eat shrimp.")
	fmt.Println(string(sSlice[len(sSlice)-le3:]))

	/*
		// 17.7.- print every letter of s with one rune per line
		for _, s := range sSlice {
			fmt.Println(string(s))
		}
	*/
	// 17.- Aux
	fmt.Println("Lengt of all", len(sSlice))

}
