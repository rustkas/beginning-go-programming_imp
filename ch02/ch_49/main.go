package main

import (
	"fmt"
)

func main() {
	/* Defining Local Variables */
	var val1 int = 50
	var val2 int = 80

	

	/* Function call to findMaximum() */
	var retVal = findMaximum(val1, val2)
	fmt.Printf("Maximum value: %d\n", retVal)
}

/* Function to find maximum value between the two numbers */
func findMaximum(num1, num2 int) int {
	/* Declaring Local Variables */
	var maxVal int
	if num1 > num2 {
		maxVal = num1
	} else {
		maxVal = num2
	}
	return maxVal
}
