
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter Some Text: ")
	input, _ := reader.ReadString('\n')
	fmt.Println("You have entered: ", input)
}
