package main

import "fmt"

func main() {
	fmt.Println("Channels as arguments and returns")
	fmt.Println("")

	myChannel := incrementador()
	cSum := puller(myChannel)

	for n := range cSum {
		fmt.Println(n)
	}
}

func incrementador() chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			out <- -i
		}
		close(out)
	}()
	return out
}

func puller(c chan int) chan int {
	out := make(chan int)
	go func() {
		var sum int
		for n := range c {
			sum += n
		}
		out <- sum
		close(out)
	}()
	return out
}
