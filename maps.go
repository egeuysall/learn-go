package main

import "fmt"

func createMap() string {
	ages := map[string]int{
		"John": 21,
		"Mary": 24,
	}
	ages["Ege"] = 15

	nums := make(map[string]int)
	nums["num"] = 24
	nums["num"] = 12
	// Here, the key will be a string and the value will be an integer.

	return fmt.Sprintf("The length will be %d!", len(nums))
}

func main() {
	println(createMap())
}
