package main

import "fmt"

func bulkSend(numMessages int) float64 {
	totalCost := 0.0
	for i := 0; i < numMessages; i++ {
		totalCost += 1.0 + (0.01 * float64(i))
	}
	return totalCost
}

func tryWhile(plantHeight int) int {
	plantHeight = 0
	for plantHeight < 10 {
		plantHeight += 1
	}
	return plantHeight
}

func infiniteLoop(num int) int {
	for {
		n := num
		n++
	}
}

func isEven(num int) string {
	if num%2 == 0 {
		return "even"
	}
	return "odd"
}

func main() {
	num := 7
	fmt.Printf("Your messages costs $%.2f! \n", bulkSend(7))
	fmt.Printf("The plant height is %d centimeters. \n", tryWhile(3))
	fmt.Printf("%d is an %v number!", num, isEven(num))
}
