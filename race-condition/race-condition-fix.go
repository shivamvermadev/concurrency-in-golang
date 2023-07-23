package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var (
	counterx int64
	wgg      sync.WaitGroup
)

func test() {
	wgg.Add(2)
	// both the go routines makes their own copy for counter variable
	go incrementCounterNew(1)
	go incrementCounterNew(2)

	wgg.Wait()
	fmt.Println("final counter : ", counter)
	fmt.Println("final counterx : ", counterx)
}

func incrementCounterNew(id int) {
	defer wgg.Done()

	for count := 0; count < 2; count++ {

		atomic.AddInt64(&counterx, 1)
		runtime.Gosched() // this function  yields the thread and give the other goroutine a chance to run. This is being done in the middle of the operation to force the scheduler to swap between the two goroutines to exaggerate the effects of the race condition
	}
}
