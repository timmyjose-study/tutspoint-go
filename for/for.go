package main

import "fmt"

func main() {
	for a := 0; a < 10; a++ {
		fmt.Printf("%d ", a)
	}

	a := 1
	b := 10

	for a < b {
		fmt.Printf("%d ", a)
		a++
	}

	nums := []int{1, 2, 3, 4, 5}

	for idx, num := range nums {
		fmt.Printf("%d: %d\n", idx, num)
	}

	var j int
	for i := 2; i < 100; i++ {
		for j = 2; j <= i/j; j++ {
			if i%j == 0 {
				break
			}
		}

		if j > i/j {
			fmt.Println(i)
		}
	}
}