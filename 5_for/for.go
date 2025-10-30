package main

import "fmt"

//for -> only construct in go for looping
func  main()  {
	// while looop
	// i := 1
	// for i <= 5 {
	// 	fmt.Println(i)
	// 	i = i +1
	// }


	// infinite loop
	// for{
	// 	fmt.Println("1")
	// }

	// classic for loop
	// for i :=0; i<= 5; i++{
	// 	// loop will stop
	// 	// break

	// 	if i == 3{
	// 		continue
	// 	}
	// 	fmt.Println(i);
	// }

	//1.22 range -> to iterate over array, slice, map, string
	for i:= range 3{
		fmt.Println(i)
	}
}