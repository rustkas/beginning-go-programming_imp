package main

import (
	"fmt"
	"sort"
)

func main() {
	states := make(map[string]string)
	states["WA"] = "Washington"
	states["NY"] = "New York"
	states["CA"] = "California"
	keys := make([]string, len(states))
	i := 0
	for key := range states {
		keys[i] = key
		i++
	}
	fmt.Println("Order of Keys Before Sorting: ", keys)
	sort.Strings(keys)
	fmt.Println("Order of Keys After Forting: ", keys)

	for i := range keys {
		fmt.Println(states[keys[i]])
		//outputs the value stored in the states map at the passed key
	}
}
