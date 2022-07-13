package main

import (
	"fmt"
	"sync"
)

func main() {
	a := []int{2, 4, 6, 8, 10} // Initial slice
	wg := sync.WaitGroup{}     // Creating WaitGroup sync primitive
	wg.Add(len(a))             // placing a limit for wg.Wait()

	// range a
	for _, val := range a {
		// launch goroutines with int input parameter to avoid use of shared memory
		go func(v int) {
			defer wg.Done()    // decrement wg.Wait() counter
			fmt.Println(v * v) // print square to stdout. Order will be a mess
		}(val) // passing val to anonymous func
	}

	wg.Wait() // wait fo goroutines to finish their job
}
