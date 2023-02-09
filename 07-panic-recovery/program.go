package main

import (
	"errors"
	"fmt"
	"log"
)

var DivideByZeroError = errors.New("divisor cannot be 0")

func main() {
	defer func() {
		if err, ok := recover().(error); ok {
			log.Println("panicked..", err)
			if errors.Is(err, DivideByZeroError) {
				fmt.Println("Do not divide by 0")
			}
		}
		fmt.Println("Thank you")
	}()

	divisor := 0
	q, r := divide(100, divisor)
	fmt.Printf("Dividing 100 by %d, quotient = %d and remainder = %d\n", divisor, q, r)

	//converting a panic into an error (using the "divideClient" function)
	/*
		q, r, err := divideClient(100, divisor)
		if err == DivideByZeroError {
			fmt.Println("Do not attempt to divide by 0")
			return
		}
		fmt.Printf("Dividing 100 by %d, quotient = %d and remainder = %d\n", divisor, q, r)
	*/
}

func divideClient(x, y int) (quotient, remainder int, err error) {
	defer func() {
		if e := recover(); e != nil {
			err = e.(error)
			return
		}
	}()
	quotient, remainder = divide(x, y)
	return
}

//3rd party library
func divide(x, y int) (quotient, remainder int) {
	fmt.Println("Calculating quotient")
	if y == 0 {
		panic(DivideByZeroError)
	}
	quotient = x / y
	fmt.Println("Calculating remainder")
	remainder = x % y
	return
}
