package main

import (
	"fmt"
	"sync"
)

func main() {
	a := []int{2, 4, 6, 8, 10} // initial slice
	sum := 0                   // init sum value
	wg := sync.WaitGroup{}     // init sync.WaitGroup
	mu := sync.Mutex{}         // init sync.Mutex

	wg.Add(len(a)) // set WG counter
	// range a
	for _, val := range a {
		// launch goroutines
		go func(v int) {
			defer wg.Done() // decrement WG counter
			mu.Lock()       // mutex ON
			sum += v * v    // concurrency-safe operation
			mu.Unlock()     // mutex OFF
		}(val) // passing val to goroutine
	}
	wg.Wait() // waiting for goroutines to do their job
	fmt.Printf("The sum of squares: %d", sum)
}
