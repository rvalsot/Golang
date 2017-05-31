package main

import "fmt"

type Square struct {
	side float64
}

type Pentagon struct {
	side    float64
	apotema float64
}

func (z Square) area() float64 {
	return z.side * z.side
}

func (z Pentagon) area() float64 {
	return z.apotema * z.side * 2.5
}

type Shape interface {
	area() float64
}

func info(z Shape) {
	fmt.Println(z)
	fmt.Println(z.area())
}

func main() {
	fmt.Println("Interfaces")
	fmt.Println("")

	s := Square{10}
	info(s)
	p := Pentagon{1, 1}
	info(p)
}
