package main

import "fmt"

func swap(a, b *int) {
	t := *a
	*a = *b
	*b = t
}

func main() {
	a, b := 1, 2

	fmt.Printf("Before swap, a: %d, b: %d\n", a, b)
	swap(&a, &b)
	fmt.Printf("After swap, a: %d, b: %d\n", a, b)

}