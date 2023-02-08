package main

import "fmt"

func main() {
	/*
		var firstName, lastName string
		fmt.Println("Enter your name:")
		fmt.Scanf("%s %s\n", &firstName, &lastName)
		fmt.Printf("Hi %s %s, Have a nice day!\n", firstName, lastName)
	*/

	var n1, n2 int
	fmt.Println("Enter the numbers")
	fmt.Scanln(&n1, &n2)
	fmt.Println(n1, n2)
}
