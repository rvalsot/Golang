package main

import "fmt"

func max(numbers ...int) (largest int) {
	fmt.Printf("%T \n", numbers)
	for _, v := range numbers {
		if v > largest {
			largest = v
		}
	}
	return
}

func main() {
	greatest := max(1, 2, 3, 4, 5, 6, 7, 7, 8, 9, 10, 2, 3, 4, 5, 20, 21, 2)
	fmt.Println("The largest number is: ", greatest)

}
