package main

import (
	"fmt"
	"strings"
)

func main() {
	greetings := []string{"hello", ", ", "world"}

	combined := strings.Join(greetings, "")
	fmt.Println(combined)

	for _, r := range combined {
		fmt.Printf("%c\n", r)
	}
}