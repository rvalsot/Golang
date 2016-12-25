package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

type Spy struct {
	person
	LicenseToKill bool
}

func main() {
	fmt.Print("Structs")
	fmt.Println("")

	person1 := person{"Juan", "Pijamas", 12}
	person2 := person{"Porfirio", "Diaz", 75}
	person3 := Spy{person1, true}

	fmt.Println("Person 1", person1)
	fmt.Println("The boss", person2.firstName, person2.lastName, "is", person2.age, "years old.")
	fmt.Printf("person.age type: %T \n", person1.age)
	fmt.Printf("Spy type: %T \n", person3)
	fmt.Println("Person 3 (spy):", person3)
}
