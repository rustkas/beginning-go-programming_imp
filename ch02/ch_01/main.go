// Listing 2-4. Ways to Declare Variables in Go
package main

import (
	"fmt"
)

func main() {
	//Explicit Declaration
	var aStringVariable string = "i am a string"
	fmt.Println(aStringVariable)
	fmt.Printf("Printing variable along text: %s \n", aStringVariable)
	var anotherStringVariable string
	fmt.Println(anotherStringVariable)
	fmt.Printf("Printing variable along text: %s \n", anotherStringVariable)
	var defaultInt int
	fmt.Println(defaultInt)
	//Implicit Declaration
	myString := "Implicit Declaration of string"
	fmt.Println(myString)

	var (
		name     string
		age      int
		location string
	)
	name = "Anatolii"
	fmt.Printf("name: %s \n", name)
	age = 42
	fmt.Printf("age: %d \n", age)
	location = "Thailand"
	fmt.Printf("location: %s \n", location)

	var (
		name1, location1 string
		age1             int = 45
		phone            int
	)

	fmt.Printf("name: %s \n", name1)
	age = 42
	fmt.Printf("age: %d \n", age1)
	location = "Thailand"
	fmt.Printf("location: %s \n", location1)
	phone = 323344
	fmt.Printf("phone: %d \n", phone)

	var (
		name2, location2, age2 = "Prince Oberyn", "Dorne", 32
	)
	fmt.Printf("%s %s %d \n", name2, location2, age2)

	const name_const, location_const, age_const = "Prince Oberyn", "Dorne", 32
}
