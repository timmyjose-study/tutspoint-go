package main

import "fmt"

func main() {
	numbers := []int{0, 1, 2, 3, 4, 5}

	for idx, num := range numbers {
		fmt.Printf("Number at index %d = %d\n", idx, num)
	}

	ch := make(chan string, 10)
	ch <- "hello"
	ch <- "world"

	close(ch)

	for s := range ch {
		fmt.Println(s)
	}
}