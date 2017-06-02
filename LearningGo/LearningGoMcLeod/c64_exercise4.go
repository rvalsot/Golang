package main

import "fmt"

func main() {
	var numbero, numbert int

	fmt.Println("Please, enter an int number")
	fmt.Scan(&numbero)

	fmt.Println("Please, enter another number")
	fmt.Scan(&numbert)

	fmt.Println("Your number's product is:", numbero*numbert)
}
