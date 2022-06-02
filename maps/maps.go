package main

import "fmt"

func main() {
	capitals := make(map[string]string, 10)

	capitals["France"] = "Paris"
	capitals["India"] = "New Delhi"
	capitals["Japan"] = "Tokyo"
	capitals["Italy"] = "Rome"
	capitals["China"] = "Beijing"

	for country, capital := range capitals {
		fmt.Printf("%v => %v\n", country, capital)
	}

	capital, ok := capitals["United States"]
	if !ok {
		fmt.Println("Could not find the capital of the U.S.A")
	} else {
		fmt.Printf("The capital of the U.S.A is %v\n", capital)
	}

	delete(capitals, "France")
	capital, ok = capitals["France"]
	if !ok {
		fmt.Println("Could not find the capital of France")
	} else {
		fmt.Printf("The capital of France is %v\n", capital)
	}
}