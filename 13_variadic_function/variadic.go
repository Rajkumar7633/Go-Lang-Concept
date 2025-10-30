package main

import "fmt"


func sum(nums ...int) int {
	total := 0
	for _, num:= range nums{
		total = total +num
	}
	return total
}

func main() {

	nums := []int{10, 20, 30}
	result := sum(nums...)
	fmt.Println("The total is:", result)
	// result := sum(1,2,3,4,5)
}