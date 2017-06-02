package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type shape interface {
	area() float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func info(s shape) {
	fmt.Println("Area: ", s.area())
}

func main() {
	fmt.Println("Value receivers")
	fmt.Println("")

	circulo := circle{4}
	info(&circulo)
}
