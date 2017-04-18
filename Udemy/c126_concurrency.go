package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup
var counter int

func init() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func main() {
	fmt.Println("Concurrency")
	fmt.Println("")
	wg.Add(20)
	// go foo()
	// go barr()
	go incrementor("Foo")
	go incrementor("Barr")
	wg.Wait()
	fmt.Println("Final Counter", counter)
}

func incrementor(s string) {
	for i := 0; i < 20; i++ {
		x := counter
		x++
		time.Sleep(time.Duration(rand.Intn(3)) * time.Millisecond)
		counter = x
		fmt.Println(s, i, "Counter:", counter)

	}
	wg.Done()
}

func foo() {
	for i := 0; i < 10; i++ {
		fmt.Println("Foo:", i)
		time.Sleep(time.Duration(3 * time.Millisecond))
	}
	wg.Done()
}

func barr() {
	for i := 0; i < 10; i++ {
		fmt.Println("Barr:", i)
		time.Sleep(time.Duration(6 * time.Millisecond))
	}
	wg.Done()
}
