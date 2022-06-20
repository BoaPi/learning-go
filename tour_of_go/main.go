package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"time"
)

// all lessons from tour of Go bundled together
func main() {
	// declare a map of lessons
	lessons := []func(){
		welcome,
		welcomeWithTime,
		randomNumber,
		mathPackage,
		mathExports,
		mathAddingUp,
		multiReturn,
		nakedReturn,
		declareCardio,
		multiDeclareAndInitialization,
		basicTypes,
		zeroValues,
	}

	runLessons(lessons)
}

// loop over lessons map and run each one with some additional logs around
func runLessons(lessons []func()) {
	for key, fn := range lessons {
		fmt.Println("---------------")
		fmt.Println("Lesson: ", key+1)
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

// using rand function of math package to generate random number
// will always be the same, if the seed is not adjusted
func randomNumber() {
	fmt.Println("Some random number: ", rand.Int())
}

// using Printf to use distinct format in logged string
// using math package and exported Sqrt() function
func mathPackage() {
	fmt.Printf("Now you have %g problems", math.Sqrt(7))
}

// only capitalized functions & values are accessible from outside the package
func mathExports() {
	fmt.Println("Here comes the number Pi:", math.Pi)
}

// function to add up two numbers for mathAddingUp()
func add(x, y int) {
	fmt.Printf("%v plus %v = %v", x, y, x+y)
}

// using declared function and pass parameters
func mathAddingUp() {
	add(31, 75)
}

// just swaps the two input strings order
func swap(x, y string) (string, string) {
	return y, x
}

// multiple returns
func multiReturn() {
	a := "Hello"
	b := "World"

	swappedA, swappedB := swap(a, b)

	fmt.Printf("Original Order: %q %q \n", a, b)
	fmt.Printf("New Order: %q %q", swappedA, swappedB)
}

// some random computations with numbers and a naked return
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x

	return
}

// use naked return of slit function
func nakedReturn() {
	fmt.Println(split(10))
}

// declare variables lesson
func declareCardio() {
	var c, python, java bool
	var i int

	fmt.Println(c, python, java, i)
}

// declare and initialize multiple variables
func multiDeclareAndInitialization() {
	i, j := 1, 2

	fmt.Println("Using add() function from former lesson here")
	add(i, j)
}

// some basic types
var (
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

// basic types examples
func basicTypes() {
	fmt.Printf("Type: %T Value: %v \n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v \n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v \n", z, z)
}

// zero values of declarations
var (
	i int
	f float64
	b bool
	s string
)

// zero value lesson
func zeroValues() {
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}
