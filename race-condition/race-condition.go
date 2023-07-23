package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int
	wg      sync.WaitGroup
)

func main() {
	wg.Add(2)
	// both the go routines makes their own copy for counter variable
	go incrementCounter(1)
	go incrementCounter(2)

	wg.Wait()
	fmt.Println("final counter : ", counter)
	test()
}

func incrementCounter(id int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		value := counter

		runtime.Gosched() // this function  yields the thread and give the other goroutine a chance to run. This is being done in the middle of the operation to force the scheduler to swap between the two goroutines to exaggerate the effects of the race condition

		value++
		counter = value
	}
}
