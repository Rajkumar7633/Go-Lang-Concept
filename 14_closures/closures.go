package main


func counter() func() int {
	var count int =0

	return func() int {
		count++
		return count
	}
}

func main() {

	increment := counter()

	println(increment()) // Output: 1
	println(increment()) // Output: 2
	println(increment()) // Output: 3	

}