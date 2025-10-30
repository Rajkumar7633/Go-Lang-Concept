package main

import "fmt"

func main() {
	// age := 10

	// if age >= 45{
	// 	fmt.Println("person is an adult")
	// } else {
	// 	fmt.Println("Person in not an adult")
	// }

	// this is another method use
	// if age >= 19 {
	// 	fmt.Println("person is and adult")
	// }else if age >= 12 {
	// 	fmt.Println("person is teenager")
	// }else{
	// 	fmt.Println("Person is kid")
	// }

	// logical operators

	// var role = "admin"
	// var hasPermissions =true
	// if role == "admin" || hasPermissions{
	// 	fmt.Println("Yes its good")
	// }else{
	// 	fmt.Println("No its denied")
	// }

	// we can declare a variable inside if construct
	if age := 20; age >= 18 {
		fmt.Println("Person is an adult", age)
	} else if age >= 12 {
		fmt.Println("person is teenager", age)
	}

	// go does not have ternary operator, you will have to use normal if else

}
