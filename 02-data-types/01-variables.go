package main

import "fmt"

/*
	var userName string = "Suresh"
*/

/*
	var userName = "Suresh"
*/

// userName := "Suresh" // Invalid. Cannot use := at the package scope

var future_var string

func main() {
	/*
		var userName string
		userName = "Magesh"
	*/

	/*
		var userName string = "Magesh"
	*/

	//type inference
	/*
		var userName = "Magesh" //userName data type is "inferred" to be a "string"
	*/

	userName := "Magesh"
	fmt.Printf("Hi %s, Have a nice day!\n", userName)

	//unused variable
	/*
		var x int
		x = 200
		fmt.Println(x)
	*/

	//multiple variables
	/*
		var x int
		var y int
		var str string
		var result int
		x = 100
		y = 200
		str = "Sum of %d and %d is %d\n"
		result = x + y
		fmt.Printf(str, result)
	*/

	/*
		var x, y, result int
		var str string
		x = 100
		y = 200
		str = "Sum of %d and %d is %d\n"
		result = x + y
		fmt.Printf(str, x, y, result)
	*/

	/*
		var (
			x, y, result int
			str          string
		)
		x = 100
		y = 200
		str = "Sum of %d and %d is %d\n"
		result = x + y
		fmt.Printf(str, x, y, result)
	*/

	/*
		var x int = 100
		var y int = 200
		var str string = "Sum of %d and %d is %d\n"
		var result int = x + y
		fmt.Printf(str, x, y, result)
	*/

	/*
		var x, y int = 100, 200
		var str string = "Sum of %d and %d is %d\n"
		var result int = x + y
		fmt.Printf(str, x, y, result)
	*/

	/*
		var (
			x, y   int    = 100, 200
			str    string = "Sum of %d and %d is %d\n"
			result int    = x + y
		)
		fmt.Printf(str, x, y, result)
	*/

	//type inference
	/*
		var (
			x, y   = 100, 200
			str    = "Sum of %d and %d is %d\n"
			result = x + y
		)
		fmt.Printf(str, x, y, result)
	*/

	/*
		var (
			x, y, str = 100, 200, "Sum of %d and %d is %d\n"
			result    = x + y
		)
		fmt.Printf(str, x, y, result)
	*/

	x, y, str := 100, 200, "Sum of %d and %d is %d\n"
	result := x + y
	fmt.Printf(str, x, y, result)

}
