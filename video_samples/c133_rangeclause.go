package main

import "fmt"

func main() {
	fmt.Println("Channels")
	fmt.Println("")

	channel := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- i
			fmt.Println("Iteration:", i)
		}
		close(channel)
	}()

	for n := range channel {
		fmt.Println(n)
	}
}
