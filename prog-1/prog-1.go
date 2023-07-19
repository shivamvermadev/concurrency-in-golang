package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)

	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("starting goroutines")

	go func() {
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'a'; char <= 'a'+26; char++ {
				fmt.Printf("%c ", char)
				time.Sleep(time.Millisecond * 1000)
			}
		}
	}()

	go func() {
		defer wg.Done()

		for count := 0; count < 3; count++ {
			for char := 'A'; char <= 'A'+26; char++ {
				fmt.Printf("%c ", char)
				time.Sleep(time.Millisecond * 1000)
			}
		}
	}()

	fmt.Println("Waiting to finish")
	wg.Wait()
	fmt.Println("\n Terminating process")
}
