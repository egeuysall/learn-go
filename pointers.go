package main

import "fmt"

func main() {
	x := 5
	y := x
	z := &y // This syntax gives us the address of the x variable in the memory.
	*z = 6

	fmt.Printf("My number is %d!", y)
}
