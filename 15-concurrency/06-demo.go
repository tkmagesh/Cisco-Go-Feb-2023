/* concurrency safe data manipulation */
package main

import (
	"fmt"
	"sync"
)

var count int
var mutex sync.Mutex

func main() {
	wg := &sync.WaitGroup{}
	for i := 0; i < 200; i++ {
		wg.Add(1)
		go increment(wg)
	}
	wg.Wait()
	fmt.Println("count :", count)
}

func increment(wg *sync.WaitGroup) {

	defer wg.Done()
	mutex.Lock()
	{
		count++
	}
	mutex.Unlock()
}
