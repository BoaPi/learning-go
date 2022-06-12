package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// returns a greeting for the passed name
func Hello(name string) (string, error) {
	// return an error if the name is empty
	if name == "" {
		return "", errors.New("empty name")
	}

	message := fmt.Sprintf(randomFormat(), name)

	return message, nil
}

// Hellos returns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)

		if err != nil {
			return nil, err
		}

		messages[name] = message

	}

	return messages, nil
}

func init() {
	rand.Seed(time.Now().Unix())
}

// returns a set of greetings, which will be randomly selected from
func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// selects on greeting form formats by using a random index
	// based on the length of formats
	return formats[rand.Intn(len(formats))]
}
