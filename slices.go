package main

import "fmt"

func getMessages() []string {
	// Arrays example
	messages := [3]string{
		"Click here",
		"Sign up",
		"Log in",
	}

	getSlice := messages[0:2]

	var parts string
	trySlice := []string{
		"Hello!",
		"How are you doing?",
	}
	parts = trySlice[0]
	value := fmt.Sprintf("My part is %v", parts)
	println(value)

	return getSlice
}

func main() {
	// Arrays example
	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Printf("The array is %v!\n", primes)
	getMessages()

	nums := []int{1, 2, 3, 4, 5, 6, 7}
	for _, v := range nums {
		fmt.Printf("\nYour number: %d ", v)
	}
}
