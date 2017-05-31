package main

import (
	"fmt"
	"math"
)

/*
 * Tasks for this hands-on exercise:
 * Create type square, type circle, type Rhombus, pentagon
 * Attach a method to each to calculate area & perimeter
 * Create a type shape which defines an interface as anything with area and perimeter methods
 * Create a func info hich  takes type shape and then prints the area
 * Create a value of type: circle, square, Rhombus, pentagon
 * Yse func to print the area of each variable
 */

// Shapes structures definition -----------------------------------------------

type square struct {
	Name       string
	SideLength float64
}

type circle struct {
	Name         string
	RadiusLength float64
}

type rhombus struct {
	Name                string
	MayorDiagonalLength float64
	MinorDiagonalLength float64
}

type pentagon struct {
	Name      string
	Perimeter float64
	Apothem   float64
}

// Areas ----------------------------------------------------------------------

func (s square) area() float64 {
	area := math.Pow(s.SideLength, 2)
	return area
}
func (c circle) area() float64 {
	area := c.RadiusLength * c.RadiusLength * math.Pi
	return area
}
func (r rhombus) area() float64 {
	area := r.MayorDiagonalLength * r.MinorDiagonalLength / 2
	return area
}
func (p pentagon) area() float64 {
	area := p.Perimeter * p.Apothem * 5 / 2
	return area
}

// Perimeters -----------------------------------------------------------------

func (s square) perimeter() float64 {
	SquarePerimeter := s.SideLength * 4
	return SquarePerimeter
}

func (c circle) perimeter() float64 {
	CirclePerimeter := c.RadiusLength * 2 * math.Pi
	return CirclePerimeter
}

func (r rhombus) perimeter() float64 {
	sideLength := math.Sqrt(math.Pow(r.MayorDiagonalLength, 2) + math.Pow(r.MinorDiagonalLength, 2))
	RhombusPerimeter := sideLength * 4
	return RhombusPerimeter
}

func (p pentagon) perimeter() float64 {
	PentagonPerimeter := p.Perimeter * 5
	return PentagonPerimeter
}

// Interface ------------------------------------------------------------------

type geometricFigure interface {
	area() float64
	perimeter() float64
}

func showArea(g geometricFigure) {
	fmt.Print(g.area(), "\n")
}

func showPerimter(g geometricFigure) {
	g.perimeter()
}

// Main -----------------------------------------------------------------------
/* Activate to use
func main() {
	fmt.Println("Welcome to geometric calculator")

	cuadrito := square{
		"cuadrito",
		5,
	}

	rombito := rhombus{
		"rombito",
		5,
		4,
	}

	fmt.Print("The area of ", cuadrito.Name, " is ")
	showArea(cuadrito)
	showPerimter(rombito)

}
*/
