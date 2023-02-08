package main

import (
	"fmt"
	"log"
)

func main() {
	/*
		add(100, 200)
		subtract(100, 200)
	*/

	/*
		log.Println("Operation started...")
		add(100, 200)
		log.Println("Operation completed!")

		log.Println("Operation started...")
		subtract(100, 200)
		log.Println("Operation completed!")
	*/

	/*
		logAdd(100, 200)
		logSubtract(100, 200)
	*/

	/*
		logOperationByName("add", 100, 200)
		logOperationByName("subtract", 100, 200)
	*/

	logOperation(add, 100, 200)
	logOperation(subtract, 100, 200)
	logOperation(func(x, y int) {
		fmt.Println("Multiply Result :", x*y)
	}, 100, 200)
}

func logOperation(op func(int, int), x, y int) {
	log.Println("Operation started...")
	op(x, y)
	log.Println("Operation completed!")
}

func logOperationByName(opName string, x, y int) {
	log.Println("Operation started...")
	switch opName {
	case "add":
		add(x, y)
	case "subtract":
		subtract(x, y)
	}
	log.Println("Operation completed!")
}

func logAdd(x, y int) {
	log.Println("Operation started...")
	add(100, 200)
	log.Println("Operation completed!")
}

func logSubtract(x, y int) {
	log.Println("Operation started...")
	subtract(100, 200)
	log.Println("Operation completed!")
}

//3rd party library
func add(x, y int) {
	fmt.Println("Add Result :", x+y)
}

func subtract(x, y int) {
	fmt.Println("Subtract Result :", x-y)
}
