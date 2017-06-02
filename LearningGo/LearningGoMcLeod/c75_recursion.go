package main

import "fmt"

func factorial(x int64) int64 {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)

}

func main() {
	var your_number int64
	fmt.Println("Introduce a number to know its factorial")
	fmt.Scan(&your_number)
	fmt.Println("Your factorial is:", factorial(your_number))
}
