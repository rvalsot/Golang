package main

import "fmt"

func main() {
	if true {
		fmt.Println("This is true")
	}

	if !false {
		fmt.Println("This is not false (!false)")
	}

	b := true
	if b {
		fmt.Println(b, "is the value of the variable b, and it's a boolean")
	}

	food := "cheesecake"
	if food == "cheesecake" {
		fmt.Println("Your food is cheesecake")
	}

}
