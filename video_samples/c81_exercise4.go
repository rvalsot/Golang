package main

import "fmt"

func main() {
	var tru bool = true
	var fal bool = false

	fmt.Println("True && false", tru && fal)

	fmt.Println((true && false) || (false && true) || !(false && !false))

}
