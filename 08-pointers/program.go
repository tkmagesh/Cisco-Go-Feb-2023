package main

import "fmt"

func main() {
	var x int
	x = 100

	var xPtr *int
	xPtr = &x //value -> address
	fmt.Println(x, xPtr)

	//derefencing = address -> value
	fmt.Println(*xPtr)

	fmt.Println(x == *(&x))

	fmt.Println("Before incrementing, x = ", x)
	increment(&x)
	fmt.Println("After incrementing, x = ", x)

	var a, b int = 100, 200
	fmt.Printf("Before swapping, a = %d and b = %d\n", a, b)
	swap( /*  */ )
	fmt.Printf("After swapping, a = %d and b = %d\n", a, b)
}

func increment(no *int) {
	fmt.Println(no)
	*no++
}

func swap( /*  */ ) {
	/*  */
}
