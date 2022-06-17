package main

import (
	"fmt"
	"time"
)

// all lessons from tour of Go bundled together
func main() {
	// declare a map of lessons
	lessons := map[int]func(){
		1: welcome,
		2: welcomeWithTime,
	}

	welcome()
}

// Welcome
func welcome() {
	fmt.Println("Welcome to the Tour of Go")
}

// Welcome with some time
func welcomeWithTime() {
	fmt.Println("Welcome to the Tour of Go")

	fmt.Println("It is: ", time.Now())
}
