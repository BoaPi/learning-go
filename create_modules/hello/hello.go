package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	// use greeting package 'Hello' function
	var message string
	message = greetings.Hello("Gladys")

	fmt.Println(message)
}
