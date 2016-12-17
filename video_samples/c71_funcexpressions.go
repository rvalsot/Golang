package main

import "fmt"

func goodbye() func() string {
	return func() string {
		return "See you guys!"
	}
}

func main() {
	/* This will throw an error, we cannot declare functions inside main()
	func greeting() {
		fmt.Println("Hi folk!")
	}
	*/
	// But we can have Anonymous Functions
	greeting := func() {
		fmt.Println("Hi folk!")
	}

	greeting()
	fmt.Printf("%T \n", greeting)

	seeya := goodbye()
	fmt.Printf("%T \n", seeya)
}
