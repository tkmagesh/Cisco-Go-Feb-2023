package main

import (
	"errors"
	"fmt"
)

var DivideByZeroError = errors.New("divisor cannot be 0")

func main() {
	divisor := 0
	q, r, err := divide(100, divisor)
	if err == DivideByZeroError {
		fmt.Println("Do not attempt to divide by 0")
		return
	}
	if err != nil {
		fmt.Println("something went wrong. ", err)
		return
	}
	fmt.Printf("Dividing 100 by %d, quotient = %d and remainder = %d\n", divisor, q, r)
}

/*
func divide(x, y int) (int, int, error) {
	if y == 0 {
		// err := fmt.Errorf("divisor cannot be %d", y)
		err := DivideByZeroError
		return 0, 0, err
	}
	quotient, remainder := x/y, x%y
	return quotient, remainder, nil
}
*/

func divide(x, y int) (quotient, remainder int, err error) {
	if y == 0 {
		err = DivideByZeroError
		return
	}
	quotient, remainder = x/y, x%y
	return
}
