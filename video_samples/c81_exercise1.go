package main

import "fmt"

func evenizer(num int) (halfed int, even bool) {
	halfed = num / 2

	if num%2 == 0 {
		even = true
	} else {
		even = false
	}

	return
}

func main() {
	var input int

	fmt.Println("Please, enter a number to be evaluated:")
	fmt.Scan(&input)

	halfed, even := evenizer(input)
	fmt.Println("(", halfed, ",", even, ")")

}
