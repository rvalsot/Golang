package main

import "fmt"

func main() {
	// Matched switch
	switch "Caribu" {
	case "Beaver":
		fmt.Println("I like making dams")
	case "Caribu":
		fmt.Println("I like eating grass")
	default:
		fmt.Println("You are not a canadian animal")
	}

	// Empty switch
	switch {
	case false:
		fmt.Println("This is false")
	case true:
		fmt.Println("This is true")
	default:
		fmt.Println("This is the default")
	}
}
