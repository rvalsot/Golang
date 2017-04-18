package main

import "fmt"

func main() {
	fmt.Println("Incrementor with channels")
	fmt.Println("")

	c1 := incrementor("Foos:")
	c2 := incrementor("Barr:")

	c3 := puller(c1)
	c4 := puller(c2)

	fmt.Println("Final counter:", <-c3+<-c4)

}

func incrementor(s string) chan int {
	out := make(chan int)
	go func() {
		for i := 0; i < 20; i++ {
			out <- 1
			fmt.Println(s, i)
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
