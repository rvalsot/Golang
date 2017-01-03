package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channels workout")
	fmt.Println("")

	aChannel := make(chan int)

	var waitGroup sync.WaitGroup
	waitGroup.Add(2)

	go func() {

		for i := 0; i < 10; i++ {
			aChannel <- i
		}
		waitGroup.Done()
	}()

	go func() {
		for j := 0; j < 10; j++ {
			aChannel <- j
		}
		waitGroup.Done()
	}()

	go func() {
		waitGroup.Wait()
		close(aChannel)
	}()

	for n := range aChannel {
		fmt.Print(n)
	}
}
