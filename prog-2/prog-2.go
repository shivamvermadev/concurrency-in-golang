package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(1)
	wg.Add(2)

	fmt.Println("Create goroutines")

	go printPrime("A")
	go printPrime("B")

	fmt.Println("Waiting to finish")
	wg.Wait()
	fmt.Println("Terminating Process")
}

func printPrime(prefix string) {
	defer wg.Done()
next:
	for i := 2; i < 5000; i++ {
		for j := 2; j < i; j++ {
			if i%j == 0 {
				continue next
			}
		}
		fmt.Printf("%s : %d\n", prefix, i)
	}
	fmt.Println("Completed ", prefix)

}
