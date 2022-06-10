package greetings

import "fmt"

// returns a greeting for the passed name
func Hello(name string) string {
	var message string
	message = fmt.Sprintf("Hi, %v. Welcome!", name)

	return message
}
