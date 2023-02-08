package main

import "fmt"

func main() {
	var i int8 = 100
	var f float32
	//use the typename like a function for type conversion
	f = float32(i)
	fmt.Println(f)

}
