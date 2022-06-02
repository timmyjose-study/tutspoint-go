package main

import "fmt"

func main() {
	var marks int
	var grade = "F"

	fmt.Scanf("%d", &marks)

	switch marks {
	case 90, 100:
		grade = "A"
	case 60, 70, 80:
		grade = "B"
	case 50:
		grade = "C"
	default:
		grade = "F"
	}

	fmt.Printf("You scored %d marks which gives you a grade of %v\n", marks, grade)

	var x interface{}

	switch x.(type) {
	case nil:
		fmt.Printf("type of x is: %T\n", nil)
	case int:
		fmt.Println("x is int")

	case float64:
		fmt.Println("x is float64")

	case func(int) float64:
		fmt.Println("x is a func(int) float64")

	default:
		fmt.Println("x is something else")

	}

}