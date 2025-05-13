package main

import (
	"errors"
	"fmt"
	"time"
)

// fn is a function that returns an error — this could be an HTTP call, DB query, etc.
func ApplyWithRetry(fn func() error, retries int, delay time.Duration) error {
	var err error
	for attempt := 0; attempt <= retries; attempt++ {
		err = fn()
		if err == nil {
			return nil
		}
		if attempt < retries {
			time.Sleep(delay)
		}
	}
	return fmt.Errorf("all retries failed: %w", err)
}

func main() {
	myFunc := func() error {
		fmt.Println("Trying...")
		// Simulate failure
		return errors.New("failed attempt")
	}

	err := ApplyWithRetry(myFunc, 3, time.Second)
	if err != nil {
		fmt.Println("Final Error:", err)
	}
}
