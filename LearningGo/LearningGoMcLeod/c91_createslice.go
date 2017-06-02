package main

import "fmt"

func main() {
	fmt.Print("Creating slices")
	fmt.Println("")
	student := []string{}
	students := [][]string{}
	fmt.Println(student)
	fmt.Println(students)
	fmt.Println(student == nil)
}
