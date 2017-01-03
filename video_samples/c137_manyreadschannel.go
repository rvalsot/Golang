package main

import "fmt"

func main() {
	fmt.Println("Class # Theme ✓")
	fmt.Println("")

	readers := 6
	myChannel := make(chan int)
	done := make(chan bool)

	go func() {
		for i := 0; i < 25; i++ {
			myChannel <- i
		}
		close(myChannel)
	}()

	for i := 0; i < readers; i++ {
		fmt.Println("Reading #", i)
		go func() {
			for readers := range myChannel {
				fmt.Println(readers)
			}
		}()
	}

	for i := 1; i < readers; i++ {
		<-done
	}

}
