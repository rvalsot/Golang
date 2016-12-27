package main

import (
	"fmt"
	"sort"
)

type people []string

func (p people) Len() int {
	return len(p)
}

func (p people) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p people) Less(i, j int) bool {
	return p[i] < p[j]
}

func main() {
	fmt.Println("Sorting")
	fmt.Println("")

	toSort := people{"Zenox", "Constantin", "Xerxes", "Augustus", "Vercingetorix"}

	fmt.Println(toSort)
	sort.Sort(toSort)
	fmt.Println(toSort)
	sort.Sort(sort.Reverse(sort.StringSlice(toSort)))
	fmt.Println("Reversed:", toSort)
	fmt.Println("----------------------------------------")

	stringSort := []string{"Grytviken", "Kerguelen", "Bouvet", "Svalbard"}

	fmt.Println(stringSort)
	sort.StringSlice(stringSort).Sort()
	fmt.Println(stringSort)
	sort.StringSlice(stringSort).Swap(1, 2)

}
