package main

import (
	"errors"
	"fmt"
)

func doSomething(bad bool) error {
	if bad {
		return errors.New("something went wrong")
	}
	return nil
}

func main() {
	err := doSomething(true)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Success")
}
