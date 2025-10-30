package main

import (
	"fmt"
	"slices"
)

// slice -> dynamic array
// most use construct in go
func main() {
	//unitilized slice in nil
	//  var nums []int

	//  fmt.Println(nums==nil)// true
	// fmt.Println(len(nums)) // 0

	// var nums = make([]int, 2, 5)

	//  capacity -> maximum numbers of elements can fit
	// fmt.Println(cap(nums)) // 5
	// fmt.Println(nums) // [0 0]

	// nums = append(nums, 1)
	// nums = append(nums, 2)
	// nums = append(nums, 3)
	// nums = append(nums, 4)
	// nums = append(nums, 5)
	// nums = append(nums, 6) // when capacity exceeded, new underlying array is created with double capacity

	// nums := []int{}
	// nums= append(nums, 1)
	// nums = append(nums, 2)


	// nums[0] = 3
	// fmt.Println(nums)
	// fmt.Println(cap(nums))
	// fmt.Println(len(nums))


	// var nums = make([]int, 0 ,5)
	// nums = append(nums, 2)
	// var nums2 = make([]int, len(nums))
	// //copy funciton
	// copy(nums2, nums)

	// fmt.Println(nums, nums2)



	//slice operator

	// var nums = []int{1,2,3}
	// // fmt.Println(nums[0:1]) // [1]
	// // fmt.Println(nums[:2])
	// fmt.Println(nums[1:])


	//slice
	var nums1= []int{1,2}
	var nums2 = []int{1,2}
	fmt.Println(slices.Equal(nums1,nums2)) // true

	//2d slice
	var grid = [][]int{
		{1,2,3},
		{4,5,6},
		{7,8,9},
	}

	fmt.Println(grid) // [[1 2 3] [4 5 6] [7 8 9]]
}
