package main

import "fmt"

func main() {
	a := 10
	fmt.Printf("%d is located at address %x\n", a, &a)

	p := &a
	*p += 100
	fmt.Println(a)

	pp := &p
	**pp += 200
	fmt.Println(a)

	var ptr int
	fmt.Println(ptr)

	nums := []int{1, 2, 3, 4, 5}
	ptrs := make([]*int, 5, 5)

	for idx := range nums {
		ptrs[idx] = &nums[idx]
	}

	for idx := range ptrs {
		fmt.Println(*ptrs[idx])
	}
}