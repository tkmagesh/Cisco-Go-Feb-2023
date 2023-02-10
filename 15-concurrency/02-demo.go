package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1) // incrementing the wg counter by 1
	go f1()   // scheduled to be executed by the scheduler
	f2()
	wg.Wait() // block until the wg counter becomes 0
}

func f1() {
	fmt.Println("f1 started")
	time.Sleep(3 * time.Second)
	fmt.Println("f1 completed")
	wg.Done() // decrementing the wg counter by 1
}

func f2() {
	fmt.Println("f2 invoked")
}
