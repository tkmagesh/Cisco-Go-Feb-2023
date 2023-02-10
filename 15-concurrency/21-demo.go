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
	ch := genFib(10 * time.Second)
	for no := range ch {
		fmt.Println(no)
	}
}

func genFib(d time.Duration) <-chan int {
	ch := make(chan int)
	timeoutCh := time.After(d)
	go func() {
		x, y := 0, 1
	LOOP:
		for {
			select {
			case <-timeoutCh:
				break LOOP
			default:
				time.Sleep(500 * time.Millisecond)
				ch <- x
				x, y = y, x+y
			}
		}
		close(ch)
	}()
	return ch
}
