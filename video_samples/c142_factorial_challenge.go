package main

import "fmt"

func main() {
	fmt.Println("Factorial Challenge")
	fmt.Println("Output the factorial of a number via a Channel")

	toFactorize := factoChannel(5)
	theFactorial := factoChannelPuller(toFactorize)

	fmt.Println("The factorial:", <-theFactorial)

}

func factoChannel(number int) chan int {
	output := make(chan int)
	go func() {
		for i := 1; i <= number; i++ {
			output <- i
		}
		close(output)
	}()
	return output
}

func factoChannelPuller(inChan chan int) chan int {
	output := make(chan int)

	go func() {
		var factorial int = 1

		for n := range inChan {
			factorial *= n
		}
		output <- factorial
		close(output)
	}()

	return output
}
