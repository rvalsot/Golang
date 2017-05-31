package main 

import "fmt"

func hi() {
	fmt.Println("Hi")
}

func secondHi(){
	fmt.Println("Second Hi")
}

func thirdHi(){
	fmt.Println("Third Hi")
}

func main () {
	defer hi()
	defer secondHi()
	thirdHi()
}
