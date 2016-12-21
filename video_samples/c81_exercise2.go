package main

import "fmt"

func halfer(n int) (float64, bool) {
	return float64(n) / 2, n%2 == 0
}

func main() {
	var input int

	fmt.Println("Please, enter a number to be evaluated:")
	fmt.Scan(&input)

	halfed, even := halfer(input)
	fmt.Println("(", halfed, ",", even, ")")

}
