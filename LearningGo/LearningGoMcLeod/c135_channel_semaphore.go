package main

import "fmt"

func main() {
	fmt.Println("Channel semaphore")
	fmt.Println("")

	aChannel := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 10; i++ {
			aChannel <- i * 2
		}
		done <- true
	}()

	go func() {
		for i := 0; i < 10; i++ {
			aChannel <- -i
		}
		done <- true
	}()

	go func() {
		<-done
		<-done
		close(aChannel)
	}()

	for n := range aChannel {
		fmt.Println(n)
	}
}
