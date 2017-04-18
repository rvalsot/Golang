// Sample from https://www.socketloop.com/tutorials/golang-how-to-count-duplicate-items-in-slice-array

package main

import (
	"fmt"
)

func dupCount(list []string) map[string]int {
	duplicate_frequency := make(map[string]int)

	for _, item := range list {
		_, exist := duplicate_frequency[item]

		if exist {
			duplicate_frequency[item]++
		} else {
			duplicate_frequency[item] = 1
		}
	}

	return duplicate_frequency
}

func main() {
	duplicate := []string{"1", "1", "2", "3", "2"}

	printslice(duplicate)
	dup_map := dup_count(duplicate)
	fmt.Println(dup_map)

	for k, v := range dup_map {
		fmt.Printf("Item: %s, Count: %d \n", k, v)
	}
}
