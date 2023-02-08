package main

import (
	"fmt"
	"log"
	"time"
)

func main() {
	/*
		add(100, 200)
		subtract(100, 200)
	*/

	//logging
	/*
		logAdd := getLogOperation(add)
		logAdd(100, 200)

		logSubtract := getLogOperation(subtract)
		logSubtract(100, 200)

		logMultiply := getLogOperation(func(i1, i2 int) {
			fmt.Println("Multiply Result :", i1*i2)
		})
		logMultiply(100, 200)
	*/

	//profiling
	/*
		profileAdd := getProfileOperation(add)
		profileAdd(100, 200)
	*/

	//profiling & logging
	/*
		logAdd := getLogOperation(add)
		profileLogAdd := getProfileOperation(logAdd)
	*/

	profileLogAdd := getProfileOperation(getLogOperation(add))
	profileLogAdd(100, 200)
}

func getProfileOperation(op func(int, int)) func(int, int) {
	return func(i1, i2 int) {
		start := time.Now()
		op(i1, i2)
		elapsed := time.Since(start)
		fmt.Println("Time Taken :", elapsed)
	}
}

func getLogOperation(op func(int, int)) func(int, int) {
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
