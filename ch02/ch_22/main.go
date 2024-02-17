package main

import (
	"fmt"
)

func main() {
	var colors = []string{"ed", "Green", "Blue"}
var marks1 []float32 /* slice of unspecified size */
var marks2 = make([]int,5) /* slice of length and capacity 5*/
	fmt.Println("colors: ", colors)
	fmt.Println("marks1: ", marks1)
	fmt.Println("marks2: ", marks2)
}
