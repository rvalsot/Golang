package main

import "fmt"

func removeDuplicates(elements []int) []int {
	encountered := map[int]bool{}
	result := []int{}

	for v := range elements {
		if encountered[elements[v]] == true {
			continue
		} else {
			encountered[elements[v]] = true
			result = append(result, elements[v])
		}
	}

	return result
}

func main() {
	fmt.Println("Program that removes duplicates with a function")

	lista := []int{1, 2, 3, 4, 5, 6, 3, 2, 2, 3, 5, 1, 2, 3, 8, 11, 13, 12, 13, 49}
	fmt.Println(lista)
	removed := removeDuplicates(lista)
	fmt.Println(removed)
}
