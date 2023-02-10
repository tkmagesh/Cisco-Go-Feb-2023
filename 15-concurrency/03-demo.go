package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	// var wg sync.WaitGroup
	wg := &sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		wg.Add(1) // incrementing the wg counter by 1
		go f1(wg) // scheduled to be executed by the scheduler
	}
	f2()
	wg.Wait() // block until the wg counter becomes 0
}

func f1(wg *sync.WaitGroup) {
	defer wg.Done() // decrementing the wg counter by 1
	fmt.Println("f1 started")
	time.Sleep(3 * time.Second)
	fmt.Println("f1 completed")
}

func f2() {
	fmt.Println("f2 invoked")
}
