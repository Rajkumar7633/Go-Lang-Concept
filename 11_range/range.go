package main

import (
	"fmt"
)

// iterating over data structure
func main(){

	// nums:= []int {6,7,8}

	// for i :=0; i< len(nums); i++ {
	// 	fmt.Println(nums[i]) //. 6 7 8
	// }


	// 
	
	// m := map[string]string{"fname": "john" , "lname": "doe"}

	// for k, v := range m{
	// 	fmt.Println(k, v)
	// }
     // unicode code print here point rune
	 // starting byte of rune
	 // assic i value to store 255  but in unicode we proceed 25+++ and many
	 //255 -> 1 byte
	for i,  c := range "golang"{
		fmt.Println(i, string(c))
	}


}
