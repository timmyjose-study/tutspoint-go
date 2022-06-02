package main

import "fmt"

func main() {
	const length int = 42
	const width int = 23
	var area int

	area = length * width

	fmt.Printf("%d\n", area)
}