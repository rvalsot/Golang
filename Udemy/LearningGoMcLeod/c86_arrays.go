package main

import "fmt"

func main() {
	var x [58]string
	// fmt.Println(x)
	// fmt.Println(len(x))
	// fmt.Println(x[42])

	for i := 65; i < 122; i++ {
		x[i-65] = string(i)
	}

	// fmt.Println(x)
	// fmt.Println(len(x))
	// fmt.Println(x[42])

	var y [256]int
	fmt.Println(len(y))
	for i := 0; i < 256; i++ {
		y[i] = i
	}
	for i, v := range y {
		fmt.Printf("%v - %T - %b \n", v, v, v)
		if i > 15 {
			break
		}
	}
}
