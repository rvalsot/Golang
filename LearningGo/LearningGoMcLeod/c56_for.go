package main

import "fmt"

func main() {
	var a int = 2
	for i := 0; i <= a; i++ {
		for j := 0; j <= a; j++ {
			fmt.Println("[", i, ",", j, "]")
		}
	}
	fmt.Println("========================================")
	i := 0
	for i < 4 {
		fmt.Println("a")
		i++
	}

	fmt.Println("========================================")
	j := 0
	for {
		j++
		if j%2 == 0 {
			continue
		}
		fmt.Println("b", j)
		if j > 6 {
			break
		}
	}

}
