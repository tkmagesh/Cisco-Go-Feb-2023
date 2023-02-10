/*
Create a "genFib()" function that generates the fibonacci series
The # of fibonacci values to be generated is passed as an argument to the function
Receive the generated numbers and print them in the main() function
IMPORTANT: Use channels for communication
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	count := 10
	ch := genFib(10)
	for i := 0; i < count; i++ {
		fmt.Println(<-ch)
	}
}

func genFib(count int) <-chan int {
	ch := make(chan int)
	go func() {
		x, y := 0, 1
		for i := 0; i < count; i++ {
			time.Sleep(500 * time.Millisecond)
			ch <- x
			x, y = y, x+y
		}
	}()
	return ch
}
