package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// setup of logger
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// use greeting package 'Hello' function
	message, err := greetings.Hello("Gladys")

	// if an error was received, print to console
	// exit program
	if err != nil {
		log.Fatal(err)
	}

	// if no error was received, print greeting to the console
	fmt.Println(message)
}
