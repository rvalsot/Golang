package main

import "fmt"

func main() {
	fmt.Println("Factorial Challenge")
	fmt.Println("Output the factorial of a number via a Channel, sample of teacher McLeod")

	toFactorize := factoChannel(5)

	// Do not forget the safety net of the range!
	for n := range toFactorize {
		fmt.Println("The factorial:", <-toFactorize)
	}
}

func factoChannel(number int) chan int {
	output := make(chan int)
	go func() {
		var total int = 1
		for i := 1; i <= number; i++ {
			total *= i
		}
		output <- total
		close(output)
	}()
	return output
}
