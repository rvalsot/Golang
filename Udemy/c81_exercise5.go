package main

import "fmt"

func foo(nums ...int) {
	fmt.Println(nums)
}

func main() {
	foo(1, 2)
	foo(1, 2, 3)
	aSlice := []int{1, 2, 3, 4, 5, 67, 3}
	foo(aSlice...)
	foo()

}
