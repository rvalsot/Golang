package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Using int slice sort")
	fmt.Println("")

	nSlice := []int{12, 3, 4, 8, 6, 5, 9, 7, 6, 6, 1}
	fmt.Println(nSlice)
	sort.Sort(sort.Reverse(sort.IntSlice(nSlice)))
	fmt.Println(nSlice)
}
