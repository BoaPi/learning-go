package greetings

import (
	"errors"
	"fmt"
)

// returns a greeting for the passed name
func Hello(name string) (string, error) {
	// return an error if the name is empty
	if name == "" {
		return "", errors.New("empty name")
	}

	var message string
	message = fmt.Sprintf("Hi, %v. Welcome!", name)

	return message, nil
}
