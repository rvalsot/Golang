package main

import "fmt"

func main() {
	/*
		mySlice := []string{"a", "n", "k", "e", "e", "r", "b"}
		fmt.Println(mySlice)
		fmt.Println(mySlice[2])
		fmt.Println(mySlice[2:4])
		fmt.Println("mySlice"[2])
	*/
	/*
		fmt.Println("=============================================")
		anSlice := make([]int, 0, 5)

		fmt.Println("My Slice:", anSlice)
		fmt.Println("My Slice's length:", len(anSlice))
		fmt.Println("My Slice's capacity:", cap(anSlice))

		for i := 0; i < 10; i++ {
			anSlice = append(anSlice, i)
			fmt.Println("Len:", len(anSlice), " Cap:", cap(anSlice), " Value:", anSlice[i])
		}
	*/
	fmt.Println("=============================================")

	buenosDias := []string{
		"Buen dia",
		"Good morning",
		"Bonjour",
		"Bongiorno",
		"Gutten morgen",
		"Bonan tagon",
		"добрый день",
	}
	for i, currenEntry := range buenosDias {
		fmt.Println(i, currenEntry)
	}
	for j := 0; j < len(buenosDias); j++ {
		fmt.Println(buenosDias[j])
	}
	fmt.Println("")
	fmt.Println(buenosDias[1:2])
}
