package main

import (
	"errors"
	"fmt"
	"math"
	"os"
)

func mySqrt(f float64) (float64, error) {
	if f < 0 {
		return 0, errors.New("cannot find the square root of a negative number")
	}
	return math.Sqrt(f), nil
}

func main() {
	if r, err := mySqrt(10.0); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	} else {
		fmt.Println(r)
	}

	if r, err := mySqrt(-10.0); err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	} else {
		fmt.Println(r)
	}
}