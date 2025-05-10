package main

import "fmt"

type user struct {
	name     string
	age      int
	messages struct {
		phoneNumber int
		message     string
	}
}

type car struct {
	make  string
	model string
}

type truck struct {
	car     // It is a shorthand for embedding. Now, you can access the fields of the car struct directly from the truck struct.
	bedSize int
}

func main() {
	myUser := user{}
	myUser.name = "John Doe"

	if myUser.name == "John Doe" {
		fmt.Printf("Your name: %v \n", myUser.name)
		fmt.Printf("Your age: %v \n", myUser.age)
		fmt.Printf("Your messages: %v \n", myUser.messages)
	} else {
		fmt.Println("User not found")
	}
}
