package main

import (
	"fmt"

	"example.com/create-packages/utils"
)

func main() {
	fmt.Println(utils.ToUpper("hello"))
	fmt.Println(utils.ToLower("BYE BYE"))
}