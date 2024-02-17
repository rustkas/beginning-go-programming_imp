package main

import (
	"fmt"
)

func main() {
	// string_Map := new(map[string]int)
	// string_Map["Marks"] = 56

	string_Map := make(map[string]int)
	string_Map["Marks"] = 56
	fmt.Println(string_Map)
}
