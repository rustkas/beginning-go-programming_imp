package main

import (
	"fmt"
)

func modifyInt(val int) int {
	return val + 5
}
func main() {
	num := 500
	fmt.Println("Value of Passed Argument Before Function Call: ", num)
	fmt.Println("Value Returned by Function call:", modifyInt(num))
	fmt.Println("Value of Passed Argument After Function Call: ", num)
}
