package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}

	for i := 0; i < len(a); i++ {
		a[i] += 100
	}

	for _, n := range a {
		fmt.Println(n)
	}

	mat := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	for _, r := range mat {
		for _, c := range r {
			fmt.Printf("%d ", c)
		}
		fmt.Println()
	}

	fmt.Println(calcAverage(a, len(a)))
}

func calcAverage(a [5]int, len int) int {
	sum := 0
	for _, n := range a {
		sum += n
	}
	return sum / len
}