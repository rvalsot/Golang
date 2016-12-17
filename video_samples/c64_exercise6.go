package main

import "fmt"

func main() {
	for i := 0; i < 25; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println(i, ", Both")
		} else if i%3 == 0 {
			fmt.Println(i, ", in 3")
		} else if i%5 == 0 {
			fmt.Println(i, ", in 5")
		} else {
			continue
		}

	}
}
