package main

import "fmt"

func main() {
	var x float64
	x = 1.2345
	fmt.Printf("%g has type %T\n", x, x)

	y := 2
	fmt.Printf("%d has type %T\n", y, y)

	var a, b, c = 3, 1.4, "foo"
	fmt.Printf("%v has type %T, %v has type %T, %v has type %T\n", a, a, b, b, c, c)
}