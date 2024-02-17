package main

import (
	"fmt"
)

func swapValues(str1, str2 string) (string, string) {
	return str2, str1
}

func main() {
	val1, val2 := swapValues("Abdullah", "Hassan")
	fmt.Println("Values After Swap: ", val1, val2)
}
