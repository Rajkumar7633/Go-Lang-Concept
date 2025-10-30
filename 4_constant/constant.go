package main

import "fmt"

// const age = 30

func main(){
	// const name= "GoLang"

	// := is not allowed with const
	// when we declare one time then we cannot change it
	// name = "Java" // this will give error
	// const age = 30

	// fmt.Println("Age is:", age)


	// give multiple constants
	const(
		port = 5000
		host= "localhost"

	)
	fmt.Println("Server is running on", host, "at port", port)

}