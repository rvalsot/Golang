package main

import "fmt"

func visit(numbers []int, callback func(int)) {
	for _, n := range numbers {
		callback(n)
	}
}

func filter(numbers []int, callback func(int) bool) []int {
	xs := []int{}
	for _, n := range numbers {
		if callback(n) {
			xs = append(xs, n)
		}
	}
	fmt.Printf("%T \n", filter)
	fmt.Printf("%T \n", callback)
	return xs
}

func main() {
	visit([]int{1, 2, 3, 4}, func(n int) {
		fmt.Println(n)
	})

	fmt.Println("=====================================")
	xs := filter([]int{1, 2, 3, 4, 5, 6}, func(n int) bool {
		return n > 3
	})
	fmt.Println(xs)
}
