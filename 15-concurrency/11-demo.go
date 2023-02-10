package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		ch <- 10
	}()
	data := <-ch
	fmt.Println(data)
}
