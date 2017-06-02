package main

import "fmt"

func wrapper() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func main() {
	incrementando := wrapper()
	fmt.Println(incrementando())
	fmt.Println(incrementando())
}
