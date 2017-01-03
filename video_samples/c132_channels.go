package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Channels")
	fmt.Println("")

	channel := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			channel <- i
			fmt.Println(channel)
		}
	}()

	go func() {
		for {
			fmt.Println(<-channel, ": channel")
		}
	}()
	time.Sleep(time.Second)
}
