package main

import "fmt"

func main() {
	n := average(32, 43, 43, 12, 65, 6, 3, 6587, 48)

	fmt.Println(n)
}

func average(sf ...float64) float64 {
	fmt.Println(sf)
	fmt.Printf("%T \n", sf)

	var total float64
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}
