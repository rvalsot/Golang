package main

import "fmt"

func main() {
	age := 44
	fmt.Println(&age)

	changeMe(&age)
	fmt.Println("======================")
	fmt.Println(&age)
	fmt.Println(age)

}

func changeMe(z *int) {
	fmt.Println("Pointer:", z)
	fmt.Println("Value:", *z)
	*z = 24

	fmt.Println("Pointer:", z)
	fmt.Println("Value:", *z)
}
