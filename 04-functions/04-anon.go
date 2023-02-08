package main

import "fmt"

/*
	Anonymous functions
		- cannot have a name
		- have to be immediately invoked
*/

func main() {
	func() {
		fmt.Println("Hi there!")
	}()

	func(userName string) {
		fmt.Printf("Hi %s, Have a nice day!\n", userName)
	}("Magesh")

	var greetMsg = func(userName string) string {
		return fmt.Sprintf("Hi %s, Have a good day!", userName)
	}("Suresh")
	fmt.Println(greetMsg)

	//use the add and divide as anonymous functions

	/*

		fmt.Println(getGreetMsg("Suresh"))
		fmt.Println(add(100, 200))
	*/
	// fmt.Println(divide(100, 7))
	/*
		q, r := divide(100, 7)
		fmt.Printf("Dividing 100 by 7, quotient = %d and remainder = %d\n", q, r)
	*/

	/* use _ in places where we had to have a variable but not use it */
	/*
		q, _ := divide(100, 7)
		fmt.Printf("Dividing 100 by 7, quotient = %d\n", q)
	*/
}

/*
func add(x int, y int) int {
	return x + y
}
*/

func add(x, y int) int {
	return x + y
}

/* > 1 return results */
/*
func divide(x, y int) (int, int) {
	quotient := x / y
	remainder := x % y
	return quotient, remainder
}
*/

/* named return result */
func divide(x, y int) (quotient, remainder int) {
	quotient = x / y
	remainder = x % y
	return
}
