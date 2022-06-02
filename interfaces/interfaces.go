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

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2.0 * math.Pi * c.radius
}

type rectangle struct {
	length, width float64
}

func (r rectangle) area() float64 {
	return r.length * r.width
}

func (r rectangle) perimeter() float64 {
	return 2.0 * (r.length + r.width)
}

func processShape(s shape) {
	fmt.Printf("area = %g, perimeter = %g\n", s.area(), s.perimeter())
}

func processShapePtr(s *shape) {
	fmt.Printf("area = %g, perimeter = %g\n", (*s).area(), (*s).perimeter())
}

func main() {
	c := circle{radius: 10.0}
	processShape(c)

	r := rectangle{length: 10.0, width: 20.0}
	processShape(r)
}