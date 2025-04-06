package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
	perimeter() float64
}
type circle struct {
	radius float64
}
type rectangle struct {
	height float64
	width  float64
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}
func (r rectangle) area() float64 {
	return r.height * r.width
}
func (r rectangle) perimeter() float64 {
	return 2 * (r.height + r.width)
}
func printArea(s shape) {
	fmt.Printf("Area = %.2f\n", s.area())
}
func printPerimeter(s shape) {
	fmt.Printf("Perimeter = %.2f\n", s.perimeter())
}
func main() {
	q := circle{radius: 24}
	t := rectangle{height: 30, width: 50}

	printArea(q)
	printPerimeter(q)

	printArea(t)
	printPerimeter(t)
}
