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
		typeConversion,
		typeInference,
		constants,
		moreConstants,
		simpleLoopToSum,
		loopWithoutOptionalStatements,
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

// type conversion lesson
func typeConversion() {
	var i int = 42
	var f float64 = float64(i)
	var u uint = uint(f)

	fmt.Printf("Type: %T, Value: %v\n", i, i)
	fmt.Printf("Type: %T, Value: %v\n", f, f)
	fmt.Printf("Type: %T, Value: %v\n", u, u)

	var x, y int = 3, 4
	var z float64 = math.Sqrt(float64(x*x + y*y))
	var a uint = uint(f)

	fmt.Printf("Type: %T, Value: %v\n", x, x)
	fmt.Printf("Type: %T, Value: %v\n", y, y)
	fmt.Printf("Type: %T, Value: %v\n", z, z)
	fmt.Printf("Type: %T, Value: %v\n", a, a)
}

// type inference lesson
func typeInference() {
	// inference to complex128
	v := 42.23 + 2i
	fmt.Printf("v is type of %T\n", v)

	// inference to float64
	w := 42.23
	fmt.Printf("w is type of %T\n", w)
}

// const declaration and assignment
func constants() {
	const Pi = 3.14
	const fuzz = "Some text that will not change"

	fmt.Printf("Happy %v", Pi)
}

func needInt(x int) int {
	return x*10 + 1
}

func needFloat(x float64) float64 {
	return x * 0.1
}

// more on constants
func moreConstants() {
	const (
		big   = 1 << 100
		small = big >> 99
	)

	fmt.Printf("Type: %T, Value: %v\n", needFloat(big), needFloat(big))
	fmt.Printf("Type: %T, Value: %v\n", needInt(small), needInt(small))
	fmt.Printf("Type: %T, Value: %v\n", needFloat(small), needFloat(small))
}

// some simple loops
func simpleLoopToSum() {
	sum := 0

	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println("Sum Value: ", sum)
}

// simple loop without init and post statement
func loopWithoutOptionalStatements() {
	sum := 1

	fmt.Println("Start Value: ", sum)

	for ; sum < 1000; {
		sum += sum
		fmt.Println("Value in loop: ", sum)
	}

	fmt.Println("Final result: ", sum)
}
