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

	// A slice of names.
	names := []string{"Gladys", "Samantha", "Darrin"}

	// calling new Hellos function for array of names to generate random greetings
	messages, err := greetings.Hellos(names)

	// if an error was received, print to console
	// exit program
	if err != nil {
		log.Fatal(err)
	}

	// if no error was received, print greeting to the console
	fmt.Println(messages)
}
