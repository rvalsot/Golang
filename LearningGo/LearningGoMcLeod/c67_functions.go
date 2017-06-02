package main

import "fmt"

func main() {
	// A function description

	greeting("Apple")
	fullName("Juan", "Pijamas")
	fmt.Println(helloFullName("Rene ", "Valsot"))
}

func greeting(name string) {
	fmt.Println("Hi,", name, "!")
}

func fullName(firstName, lastName string) {
	fmt.Println("My whole name is:", firstName, lastName)
}

func helloFullName(fName, lName string) string {
	return fmt.Sprint(fName, lName)
}
