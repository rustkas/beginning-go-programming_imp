package main

import (
	"fmt"
	"math/rand"
)

func main() {

	var result string
	switch dow := rand.Intn(7) + 1; dow {
	case 1:
		result = "It's Sunday"
	case 2:
		result = "It's Monday"
	case 3:
		result = "It's Tuesday"
	case 4:
		result = "It's Thursday"
	case 5:
		result = "It's Fryday"
	case 6:
		result = "It's Satarday"
	default:
		result = "It's some other day"
	}
	fmt.Println(result)
}
