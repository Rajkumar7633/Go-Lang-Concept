package main

// import "fmt"

// func  add(a int, b int)int {
// 	return a + b
// }

// func getLanguage()(string, string, string){
// 	return "golang", "javascript", "c"
// }


// func processIt(fn func(a int)int){
// 	fn(1)
// }

func processIt() func (a int)int  {
	return func(a int) int {
		return 3
	}
}

func main(){

    // fn := func (a int) int {
	// 	return 2
	// }
	fn := processIt()
	fn(6)

	// lang1, lang2, lang3 := getLanguage()
	// fmt.Println(lang1,lang2, lang3)
	// result := add(3, 5)

	// fmt.Println(result)
}
