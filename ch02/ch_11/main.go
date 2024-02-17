package main

import (
	"fmt"
)

func main() {
	var intPointer *int /* pointer to an integer variable */
	fmt.Println("Value of intPointer: ", *intPointer)
}
