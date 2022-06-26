# learning-go 🎓
> **Disclaimer** repo which is used to learn go, by following the official documentation. No real use here other than to see the progress and stay motivated 🥳

## content of this repo

| #   | Topic                                                                                                        |
| --- | ------------------------------------------------------------------------------------------------------------ |
| 1   | [hello world](hello) with external module for a fancy print on console 🙃                                     |
| 2   | [create modules](create_modules) how to work with multiple modules                                           |
| 3   | [create packages](create_packages) how to work with multiple packages                                        |
| 4   | [web server](web_server) creating a first web server with `gin` framework and some API endpoints             |
| 5   | [Tour of Go](tour_of_go) just the complete TOur of Go, self written down to to get the syntax in the fingers |

## learned so far

* [x] what are go modules
* [x] how to init go modules
* [x] how to add external go modules by using `go mod tidy`
* [x] using `go fmt` for formatting go code
* [x] how to run go code with `go run .`
* [x] Go code is grouped in packages
* [x] packages are grouped in modules
* [x] modules specifies dependencies, Go version and other modules
* [x] how to write a package
* [x] that capitalized function names, indicates that the function can be called from an other package. This is known as `exported names`
* [x] what a format string is
* [x] how to use local modules
* [x] how to work with multiple modules by using `go.work`
* [x] how to initialize a workspace with `go work init`
* [x] how to update a workspace with `go work use`
* [x] how to add error handling
* [x] introduction to the standard error package `errors`
* [x] multiple return value
* [x] how to use `nil`
* [x] `nil` is a return value for an error, that indicated that everything is fine
* [x] how to short declare and assign variables
* [x] a `slice` is an `array` with dynamic size
* [x] an `array` has a fixed size
* [x] a slice is specified by to indices, `lower` and `higher` bound
* [x] `array[1:4]` is a half open range and will select 1st, 2nd and 3rd element
* [x] a `slice` does not hold data, it just describes a section of an underlying array
* [x] changing data in a `slice` will change the corresponding elements in the underlying array and other `slices` which have the same elements
* [x] lower case name of functions, makes them only accessible to the code in its own package
* [x] use `math/rand` to create a seed
* [x] use `init()` function to generate a new seed each program run
* [x] how to work with a slice
* [x] get to know the packages `math` and `time`
* [x] how to re-use an existing function adn build on it an other function that handles same input and output but as maps
* [x] what a `map` is, how to create a `map`, how to add entries to a `map` and what types needs to be defined on a `map`
* [x] how to declare and assign values to a `map`
* [x] how to write a simple loop
* [x] how to add test cases with the name of the package `<package_name>_test.go`
* [x] use the Go standard testing package for logging and reporting
* [x] about type `T` to pass into test functions for manage test state and formatted test logs
* [x] how to name test cases
* [x] how to assert for a result
* [x] to use `Fatalf` method to print to console and end execution
* [x] how to use `go build` command to build the binary
* [x] how to get install path for the binary with `go list -f '{{.Target}}'`
* [x] how to adjust zsh config to use installed binaries
* [x] to name a module `main.go` if it is standalone program
* [x] how to declare a `struct` for representing complex objects
* [x] `struct` is a collection of fields
* [x] how to use `pointers` and `struct` together
* [x] how to declare `struts` in different ways
* [x] how to access fields in `structs` with dot notation
* [x] to use `json:"<key-name>"` to prevent capitalized JSON keys after serialization
* [x] to use `Context.IndentedJSON` from `gin` for easy debugging, instead of `Context.JSON`
* [x] how to initialize a `gin` router
* [x] how to add a handler to a `gin` router
* [x] how to attache the `gin` router to a server and start tshe server
* [x] how to use `cURL` to fetch data from the given endpoint
* [x] pointer syntax - `*<pointer-name>`
* [x] how to see the memory address of a pointer - `&<pointer-name>`
* [x] how to receive the parameter of a URL using gins `c.Param()`
* [x] how to send a `404 - Not Found` response
* [x] short assignments are only available inside function
* [x] some basic types
* [x] zero values of certain types when declaring without assigning a value
* [x] what type inference is
* [x] `const` can not use short assignment `:=`
* [x] numeric `const` are high-precision values
* [x] untyped constants, takes the needed by its context
* [x] that there is left shift `<<` & `>>` right shift operators
* [x] `for loop` is build of `init statement`, `condition statement` and `post statement`
* [x] `init statement` and `post statement` is optional
* [x] `for loop` is Go's `while` by omitting the optional statements
* [x] a `for loop` without any statement is an infinite loop
* [x] `if statements` do not need `()` around the statement but `{}` for the body
* [x] `if statements` can have scoped variables
* [x] using `switch statements` for a sequence of `if - else statements`
* [x] how to get the running operating system with `runtime.GOOS`
* [x] `switch cases` can run functions
* [x] how to use `switch statements` for long if-then-else statements
* [x] how to use `defer` to execute a function later
* [x] that deferred functions are stacked and will be called in `last-in-first-out` order
* [x] pointers holds memory address of a value
* [x] zero value of a pointer is `nil`
* [x] to call the an address of a value with `&<value>` and store it in an pointer
* [x] how to read a value through a pointer by dereferencing or indirecting
* [x] Go has no pointer arithmetic
* [ ] more to come...
