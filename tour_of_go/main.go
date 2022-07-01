package main

import (
	"fmt"
	"math"
	"math/cmplx"
	"math/rand"
	"runtime"
	"strings"
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
		while,
		ifLesson,
		scopedIf,
		useSqrt,
		switchSimple,
		switchDay,
		switchFunctions,
		cleanSwitch,
		deferFunction,
		deferringFunctions,
		startingPointer,
		structLesson,
		pointyStruct,
		structDeclaration,
		startingArrays,
		startingSlices,
		complexSlices,
		literalSliceCardio,
		slicingSlices,
		lengthAndCapacityOfSlice,
		nilSlice,
		makeSlice,
		slicesOfSlice,
		appendSlice,
		firstRange,
		skippingRange,
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
func while() {
	sum := 1

	fmt.Println("Start Value: ", sum)

	for sum < 1000 {
		sum += sum
		fmt.Println("Value in loop: ", sum)
	}

	fmt.Println("Final result: ", sum)
}

// some if statement cardio
func sqrt(x float64) string {
	if x < 0 {
		return sqrt(-x) + "i"
	}

	return fmt.Sprint(math.Sqrt(x))
}

func ifLesson() {
	fmt.Println(sqrt(-9), sqrt(-4))
}

// scoped variables in if statement
func pow(x, n, lim float64) float64 {
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}

	return lim
}

func scopedIf() {
	fmt.Println(pow(3, 2, 10))
	fmt.Println(pow(3, 3, 20))
}

// exercise Loops and Functions, rebuild square root function
func Sqrt(x float64) float64 {
	// declare a start value for z
	fmt.Println("Run custom Sqrt function for: ", x)
	z := float64(1)
	tmp := float64(0)

	for i := 0; i <= 10; i++ {
		z -= (z*z - x) / (2 * z)

		if tmp != z {
			tmp = z
			fmt.Printf("run %v - narrowing z: %v\n", i+1, z)
		} else {
			return z
		}
	}

	return z
}

func useSqrt() {
	fmt.Printf("for 2 final value of z is: %v\n\n", Sqrt(2))
	fmt.Printf("for 3 final value of z is: %v\n\n", Sqrt(3))
	fmt.Printf("for 4 final value of z is: %v\n\n", Sqrt(4))
	fmt.Printf("for 9 final value of z is: %v\n\n", Sqrt(9))
	fmt.Printf("for 16 final value of z is: %v\n\n", Sqrt(16))
	fmt.Printf("for 20 final value of z is: %v\n\n", Sqrt(20))
}

// switch lesson
func switchSimple() {
	fmt.Print("Go run on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		println("OS X")
	case "linux":
		println("Linux")
	default:
		// all the other freebsd, openbsd, windows, plan9
		println("%s. \n", os)
	}
}

// switch case cardio
func switchDay() {
	fmt.Println("When is Saturday?")

	today := time.Now().Weekday()

	switch time.Saturday {
	case today + 0:
		println("Today - Yeah")
	case today + 1:
		println("Tomorrow - whoop weekend ahead")
	case today + 2:
		println("In 2 days - almost there")
	default:
		println("Way to far away")
	}
}

// switch statement with functions
func switchFunctions() {
	switch i := 3; i {
	case int(math.Sqrt(9)):
		fmt.Println("Nailed it")
	case int(math.Sqrt(16)):
		fmt.Println("4 it is")
	default:
		println("It is not 3 or 4!")
	}
}

// using switch as a clean way for long if-then-else statements
func cleanSwitch() {
	t := time.Now()
	fmt.Println("It is: ", t)

	switch {
	case t.Hour() < 12:
		fmt.Println("Gooooood morning!")

	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}

// deferring a function
func deferFunction() {
	defer fmt.Println("world")

	fmt.Print("Hello ")
}

// defer cardio
func deferringFunctions() {
	fmt.Println("Counting down")
	defer fmt.Println("Go Go Go")

	for i := 0; i <= 5; i++ {
		defer fmt.Println(i)
	}
}

// some pointer stuff
func startingPointer() {
	// pointer to an integer
	var p *int

	i := 42
	p = &i

	fmt.Println("Value i: ", i)
	fmt.Println("Address of i, p: ", p)
	fmt.Println("Value of pointer p: ", *p)
}

type Vertex struct {
	X int
	Y int
}

// struct
func structLesson() {

	v := Vertex{1, 10}

	fmt.Println("First struct ", v)

	fmt.Println("Vertex X: ", v.X)
	fmt.Println("Vertex Y: ", v.Y)
}

// pointer and struct
func pointyStruct() {
	var p *Vertex
	v := Vertex{12, 15}

	p = &v

	fmt.Println("Pointer p to Vertex v, Value X ", p.X)
	fmt.Println("Pointer p to Vertex v, Address p ", &p)

	p.X = 1e12

	fmt.Println("New Value X of Vertex v, after using pointer p to overwrite ", v)
	fmt.Println("Pointer p to Vertex v, Address p ", &p)
}

// declare some more struct
func structDeclaration() {
	var (
		v1 = Vertex{1, 2}
		v2 = Vertex{X: 3}
		v3 = Vertex{}
		p  = &Vertex{4, 5}
	)

	fmt.Println(v1, v2, v3, p)
}

// some array stuff
func startingArrays() {
	var a [2]string

	a[0] = "Hello"
	a[1] = "World!"

	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}
	fmt.Println("all primes until 13", primes)
}

// slice lesson
func startingSlices() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]

	fmt.Println("Slice of primes array: ", s)
}

// complex slice
func complexSlices() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}

	fmt.Printf("The Beatles are %q\n", names)

	a := names[0:2]
	b := names[1:3]

	fmt.Println("Some Beatles are: ", a)
	fmt.Println("Some Beatles are: ", b)

	b[0] = "XXX"

	fmt.Println("Some Beatles are no longer known: ", a)
	fmt.Println("Some Beatles are no longer known: ", b)
	fmt.Printf("The Beatles are %q\n", names)
}

// slice literals cardio
func literalSliceCardio() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("literally a slice ", q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println("another slice ", r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{4, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println("some complex struct slice ", s)
}

// more on slices
func slicingSlices() {
	s := []int{2, 3, 5, 7, 11, 13, 17}
	fmt.Println("original slice ", s)

	s = s[:]
	fmt.Println("sliced slice ", s)

	s = s[1:4]
	fmt.Println("sliced slice ", s)

	s = s[:2]
	fmt.Println("sliced slice ", s)

	s = s[1:]
	fmt.Println("sliced slice ", s)
}

// length and capacity of slices
func printSlice(s string, x []int) {
	fmt.Printf("%s len: %d, cap: %d %v\n", s, len(x), cap(x), x)
}

func lengthAndCapacityOfSlice() {
	s := []int{2, 3, 5, 7, 11, 13, 17}
	printSlice("s", s)

	// zero length slice
	s = s[:0]
	printSlice("s", s)

	// extending the slice again
	s = s[:4]
	printSlice("s", s)

	// drop first two values
	s = s[2:]
	printSlice("s", s)
}

// zero value of a slice is nil
func nilSlice() {
	var s []int

	fmt.Println(s, len(s), cap(s))

	if s == nil {
		fmt.Println("nil!")
	}
}

// make a slice
func makeSlice() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

// slices of slice
func slicesOfSlice() {
	// create a tic-tac-toe board
	board := [][]string{
		{"-", "-", "-"},
		{"-", "-", "-"},
		{"-", "-", "-"},
	}

	// player turns
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}

// appending some items to a slice
func appendSlice() {
	s := make([]int, 0, 2)
	printSlice("s", s)

	// works on nil slices
	s = append(s, 1)
	printSlice("s", s)

	s = append(s, 2)
	printSlice("s", s)

	// also works with multiple elements
	s = append(s, 3, 5, 7, 11)
	printSlice("s", s)
}

// range lesson
func firstRange() {
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}

// skipping index or value of range
func skippingRange() {
	pow := make([]int, 10)

	for i := range pow {
		pow[i] = 1 << uint(i)
	}

	for _, v := range pow {
		fmt.Printf("%d\n", v)
	}
}
