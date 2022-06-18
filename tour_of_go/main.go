package main

import (
	"fmt"
	"math/rand"
	"time"
)

// all lessons from tour of Go bundled together
func main() {
	// declare a map of lessons
	lessons := map[int]func(){
		1: welcome,
		2: welcomeWithTime,
		3: randomLesson,
	}

	runLessons(lessons)
}

// loop over lessons map and run each one with some additional logs around
func runLessons(lessons map[int]func()) {
	for key, fn := range lessons {
		fmt.Println("---------------")
		fmt.Println("Lesson: ", key)
		fmt.Println("")

		// call lesson code
		fn()

		fmt.Println("")
		fmt.Println("---------------")
		fmt.Println("")
	}
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

func randomLesson() {
	fmt.Println("Some random number: ", rand.Int())
}
