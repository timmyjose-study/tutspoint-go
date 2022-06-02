package main

import "fmt"

func sliceMeta(s []int) {
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(s), cap(s), s)
}

func main() {
	slice1 := []int{}
	sliceMeta(slice1)

	slice2 := make([]int, 5, 8)
	sliceMeta(slice2)

	var slice3 []int
	sliceMeta(slice3)

	slice3 = append(slice3, 100)
	sliceMeta(slice3)

	slice4 := make([]int, len(slice3), cap(slice3)*2+1)
	copy(slice4, slice3)
	sliceMeta(slice4)

}