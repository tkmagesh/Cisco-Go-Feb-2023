package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("	deferred - main")
	}()
	fmt.Println("main started")
	f1()
	fmt.Println("main completed")
}

func f1() {
	/*
		defer func() {
			fmt.Println("	deferred[1] - f1")
		}()
		defer func() {
			fmt.Println("	deferred[2] - f1")
		}()
		defer func() {
			fmt.Println("	deferred[3] - f1")
		}()
	*/

	defer fmt.Println("	deferred[1] - f1")
	defer fmt.Println("	deferred[2] - f1")
	defer fmt.Println("	deferred[3] - f1")

	fmt.Println("f1 started")
	f2()
	fmt.Println("f1 completed")
}

func f2() {
	defer func() {
		fmt.Println("	deferred - f2")
	}()
	fmt.Println("f2 invoked")
}
