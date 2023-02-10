package main

import (
	"fmt"
)

func main() {
	ch := make(chan int)
	go fn(ch)
	/*
		fmt.Println(<-ch)
		fmt.Println(<-ch)
		fmt.Println(<-ch)
		fmt.Println(<-ch)
		fmt.Println(<-ch)
	*/
	for i := 1; i <= 5; i++ {
		fmt.Println(<-ch)
	}
}

func fn(ch chan int) {
	/*
		time.Sleep(500 * time.Millisecond)
		ch <- 10
		time.Sleep(500 * time.Millisecond)
		ch <- 20
		time.Sleep(500 * time.Millisecond)
		ch <- 30
		time.Sleep(500 * time.Millisecond)
		ch <- 40
		time.Sleep(500 * time.Millisecond)
		ch <- 50
	*/
	for i := 1; i <= 5; i++ {
		// time.Sleep(500 * time.Millisecond)
		ch <- i * 10
	}
}
