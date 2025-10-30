package main

import "fmt"

// maps -> hash, objects, dictionaries
func main() {
	// This is a placeholder for the main function.

	//creating map

	// m := make(map[string]string)

	// //setting an elements
	// m["name"] = "golang"
	// m["area"] = "backend"


	//get an elements
	// fmt.Println(m["name"],m["area"])

	// IMP: if key does ot exists in the map then its return zero


	// m := make(map[string]int)
	// m["age"] = 30
	// m["price"] = 50

	// fmt.Println(len(m))
	// fmt.Println(m["phone"]) // there is empty value return vaue coming


	// delete(m, "price")
	// fmt.Println(m) //map[age:30]

	// clear(m)
	// fmt.Println(m) // empty


	m := map[string]int{"price": 40, "phones" : 3}

	// fmt.Println(m)


	v, ok := m["phones"]
	fmt.Println(v)

	if ok{
		fmt.Println("all ok")
	}else{
		fmt.Println("not ok")
	}

}
