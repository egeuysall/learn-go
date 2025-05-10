package main

import "fmt"

type users struct {
	username string
	password string
	email    string
}

func (u users) greet() string {
	return fmt.Sprintf("Hello %s!", u.username)
}

func sayHello(u users) {
	fmt.Println(u.greet())
}

func main() {
	sayHello(users{"John", "Doe", "<EMAIL>"})
}
