package main

import (
	"fmt"
	"time"
)

func main() {
	go f1() // scheduled to be executed by the scheduler
	f2()

	//DO NOT DO THIS
	time.Sleep(time.Second) // blocked for a second, and there by giving the opporturnity for the scheduler to execute the scheduled goroutines

	// fmt.Scanln()
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(3 * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
