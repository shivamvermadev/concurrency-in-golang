package main

import (
	"fmt"
	"runtime"
	"sync"
)

var (
	counter int
	wg      sync.WaitGroup
	mutex   sync.Mutex
)

func main() {

	wg.Add(2)

	go incrementCounter(1)
	go incrementCounter(2)

	wg.Wait()
	fmt.Println("final counter ", counter)
}

func incrementCounter(x int) {
	defer wg.Done()

	for count := 0; count < 2; count++ {
		mutex.Lock()
		{
			value := counter
			runtime.Gosched()
			value++
			counter = value
		}
		mutex.Unlock()
	}

}
