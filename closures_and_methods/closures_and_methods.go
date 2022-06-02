package main

import (
	"fmt"
	"math"
)

func linum() func() {
	linum := 0

	return func() {
		linum++
		fmt.Printf("Line number: %d\n", linum)
	}
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

func main() {
	getSqrRoot := func(f float64) float64 {
		return math.Sqrt(f)
	}

	fmt.Printf("%g\n", getSqrRoot(1.2345))

	l1 := linum()
	for i := 0; i < 10; i++ {
		l1()
	}

	l1 = linum()
	for i := 0; i < 5; i++ {
		l1()
	}

	c := circle{radius: 10}
	fmt.Printf("%.2f, %.2f\n", c.area(), c.perimeter())
}