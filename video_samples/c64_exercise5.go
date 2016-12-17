package main

import "fmt"

func main() {
	for i := 0; i < 13; i++ {
		if i%2 != 0 {
			fmt.Println(i)
		}
	}
}
