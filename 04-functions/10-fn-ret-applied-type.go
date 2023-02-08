package main

import (
	"fmt"
	"log"
	"time"
)

type Operation func(int, int)

func main() {
	profileLogAdd := getProfileOperation(getLogOperation(add))
	profileLogAdd(100, 200)
}

func getProfileOperation(op Operation) Operation {
	return func(i1, i2 int) {
		start := time.Now()
		op(i1, i2)
		elapsed := time.Since(start)
		fmt.Println("Time Taken :", elapsed)
	}
}

func getLogOperation(op Operation) Operation {
	return func(x, y int) {
		log.Println("Operation started...")
		op(x, y)
		log.Println("Operation completed!")
	}
}

//3rd party library
func add(x, y int) {
	fmt.Println("Add Result :", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract Result :", x-y)
}
