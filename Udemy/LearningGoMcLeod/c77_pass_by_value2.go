package main

import "fmt"

func main() {
	str := make([]string, 1, 25)
	fmt.Println(str)

	changeMe(str)
	fmt.Println("=====================")
	fmt.Println(str)

}

func changeMe(z []string) {
	z[0] = "Vaquita"
	fmt.Println(z, "in func")
}
