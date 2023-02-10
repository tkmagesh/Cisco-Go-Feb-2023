/* concurrency safe data manipulation */
package main

import (
	"fmt"
	"sync"
)

type Counter struct {
	sync.Mutex
	count int
}

func (c *Counter) inc() {
	c.Lock()
	{
		c.count++
	}
	c.Unlock()
}

var counter Counter

func main() {

	wg := &sync.WaitGroup{}
	for i := 0; i < 200; i++ {
		wg.Add(1)
		// go increment(wg)
		go func() {
			counter.inc()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("count :", counter.count)
}

/*
func increment(wg *sync.WaitGroup) {
	defer wg.Done()
	counter.inc()
}
*/
