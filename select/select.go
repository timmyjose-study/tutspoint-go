package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Inside func1")
		c1 <- 100
	}()

	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("Inside func2")
		c2 <- 200
	}()

	for i := 0; i < 2; i++ {
		select {
		case n := <-c1:
			fmt.Println("Waiting on c1")
			fmt.Printf("Got %d on channel 1\n", n)

		case m := <-c2:
			fmt.Println("Waiting on c2")
			fmt.Printf("Got %d on channel 2\n", m)

		case <-time.After(5 * time.Second):
			break

		}
	}

}