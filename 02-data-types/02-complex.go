package main

import "fmt"

func main() {
	/*
		var c1 complex64
		c1 = 4 + 5i
	*/
	c1 := 4 + 5i
	fmt.Println(c1)
	c2 := 7 + 11i
	c3 := c1 + c2
	fmt.Println(c3)

	fmt.Printf("Real : %v\n", real(c3))
	fmt.Printf("Imaginary : %v\n", imag(c3))
}
