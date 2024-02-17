package main

import (
	"fmt"
	"time"
	)
	func main() {
		parsedTime, _ := time.Parse(time.ANSIC, "Tue Nov 10 23:00:00 2009")
		fmt.Printf("The type of parsedTime is %T \n", parsedTime)
		fmt.Println(parsedTime.Format(time.ANSIC))
}
