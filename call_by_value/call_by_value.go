package main

import "fmt"

func main() {
	a, b := 1, 2

	fmt.Printf("Before swap (call by value), a: %d, b: %d\n", a, b)

	swap(a, b)

	fmt.Printf("After swap (call by value), a: %d, b: %d\n", a, b)
}

func swap(a, b int) {
	t := a
	a = b
	b = t
}