package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(1 * time.Second)
	quit := make(chan struct{})

	go func() {
		time := 0
		for {
			select {
			case <-ticker.C:
				fmt.Println("Cow + ", time)
			case <-quit:
				ticker.Stop()
				return
			}
			time++
		}

	}()

	fmt.Println("Passed")

}
