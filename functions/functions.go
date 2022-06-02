package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func max[T constraints.Ordered](n, m T) T {
	if n >= m {
		return n
	}
	return m
}

func swap[T any](x, y T) (T, T) {
	return y, x
}

func main() {
	fmt.Printf("max(%d, %d) = %d\n", 2, 3, max(2, 3))
	fmt.Printf("max(%g, %g) = %g\n", 1.23, 2.89, max(1.23, 2.89))
	fmt.Printf("max(%s, %s) = %s\n", "hello", "world", max("hello", "world"))

	a, b := swap(1, 2)
	fmt.Printf("%d %d\n", a, b)

	s1, s2 := swap("hello", "world")
	fmt.Printf("%q %q\n", s1, s2)

}