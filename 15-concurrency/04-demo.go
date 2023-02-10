package main

import (
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	wg := &sync.WaitGroup{}
	rand.Seed(10)
	var goroutinesCount int
	flag.IntVar(&goroutinesCount, "count", 0, "Number of goroutines to spin up")
	flag.Parse()
	fmt.Printf("Spinning up %d goroutines. Hit ENTER to start...", goroutinesCount)
	fmt.Scanln()
	for i := 1; i <= goroutinesCount; i++ {
		wg.Add(1)    // incrementing the wg counter by 1
		go fn(i, wg) // scheduled to be executed by the scheduler
	}
	wg.Wait() // block until the wg counter becomes 0
	fmt.Println("Hit ENTER to shutdown...")
	fmt.Scanln()
}

func fn(id int, wg *sync.WaitGroup) {
	defer wg.Done() // decrementing the wg counter by 1
	fmt.Printf("fn[%d] started\n", id)
	time.Sleep(time.Duration(rand.Intn(20)) * time.Second)
	fmt.Printf("fn[%d] completed\n", id)
}
