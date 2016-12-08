/*
 * Also includes class 48
 */

package main

import "fmt"

const p string = "Taxation is theft"

func main() {
	const q = 42
	fmt.Println("P -> ", p)
	fmt.Println("Q -> ", q)
	fmt.Print("")

	const (
		a = iota
		b
		c
	)
	fmt.Println("const a: ", a, " const b: ", b, " const c: ", c)

	x := 42
	fmt.Printf("%T \n", x)
}
