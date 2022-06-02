package main

import "fmt"

func main() {
	var a int

	a = 20
	if a < 20 {
		fmt.Printf("%d is less than 20\n", a)
	} else {
		fmt.Printf("%d is greater than or equal to 20\n", a)
	}

	b := 20

	if b == 10 {
		fmt.Println("equal to 10")
	} else if b == 20 {
		fmt.Println("equal to 20")
	} else if b == 30 {
		fmt.Println("equal to 30")
	} else {
		fmt.Println("Something else")
	}

	x := 100
	y := 200

	if x == 100 {
		if y == 200 {
			fmt.Println("x is 100 and y is 200")
		}
	}
}