package main

import "fmt"

func main() {
	var sayHi func()
	/*
		sayHi = func() {
			fmt.Println("Hi there!")
		}
	*/
	if sayHi == nil {
		fmt.Println("sayHi is not a function")
	} else {
		sayHi()
	}

	var greet func(string)
	greet = func(userName string) {
		fmt.Printf("Hi %s, Have a nice day!\n", userName)
	}
	greet("Magesh")

	var getGreetMsg func(string) string
	getGreetMsg = func(userName string) string {
		return fmt.Sprintf("Hi %s, Have a good day!", userName)
	}
	fmt.Println(getGreetMsg("Suresh"))

	var operation func(int, int) int

	//add operation
	operation = func(x, y int) int {
		return x + y
	}
	fmt.Println(operation(100, 200))

	//multiply operation
	operation = func(x, y int) int {
		return x * y
	}
	fmt.Println(operation(100, 200))

	var divide func(int, int) (int, int)
	divide = func(x, y int) (quotient, remainder int) {
		quotient = x / y
		remainder = x % y
		return
	}
	q, _ := divide(100, 7)
	fmt.Printf("Dividing 100 by 7, quotient = %d\n", q)
}
