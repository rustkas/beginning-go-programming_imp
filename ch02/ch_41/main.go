package main

import (
	"fmt"
)

func main() {
	total := 0
	for k := 0; k < 5; k++ {
		total += k
	}
	nums := []int{2, 3, 4}
	total = 0
	for _, num := range nums {
		total += num
	}
	fmt.Println("Total:", total)
	
	for j, num := range nums {
		if num == 2 {
			fmt.Println("Index:", j)
		}
	}
	for i := range nums {
		fmt.Println(nums[i])
	}
}
