/* concurrency safe data manipulation */
package main

import (
	"fmt"
	"runtime/debug"
	"sync"
)

var count int

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 1; i++ {
		wg.Add(1)
		go increment(wg)
	}
	wg.Wait()
	fmt.Println("count :", count)
}

func increment(wg *sync.WaitGroup) {
	debug.PrintStack()
	defer wg.Done()
	count++
}
