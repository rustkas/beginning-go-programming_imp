package main

import (
	"fmt"
)

func main() {
	var anInt int = 5
var aFloat float32 = 42
// sum := anInt + aFloat //invalid operation
sum := float32(anInt)  + aFloat
fmt.Printf("Sum: %v, Type: %T \n", sum, sum)
}
