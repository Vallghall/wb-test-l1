package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/signal"
	"time"
)

func main() {
	ch := make(chan string)                                 // init chan
	n := 3                                                  // set number of workers
	ctx, cancel := context.WithCancel(context.Background()) // init context to finish workers gracefully
	sig := make(chan os.Signal, 1)                          // init channel waiting for interrupt signal
	signal.Notify(sig, os.Interrupt)                        // listen to os.Interrupt signal
	runWorkers(ctx, ch, n)                                  // run workers

	r := bufio.NewScanner(os.Stdin) // init stdin reader
	// consistent read of data from stdin and sending it to ch
	for {
		select {
		// if os.Interrupt is received
		case <-sig:
			cancel()                // cancel workers with context
			time.Sleep(time.Second) // wait for workers to finish
			return
		default:
			if r.Scan() {
				ch <- r.Text() // send data to ch
			}
		}
	}
}

// runWorkers launches N goroutines that read values from chan and print them to stdout
func runWorkers(ctx context.Context, ch chan string, n int) {
	for i := 0; i < n; i++ {
		// launch goroutine
		go func(ind int) {
			// make goroutine listen to ch and print out values from it
			for {
				select {
				// on cancelFunc call when os.Interrupt is received
				case <-ctx.Done():
					fmt.Printf("finish worker %d\n", ind)
					return
				case val := <-ch:
					fmt.Printf("value %s from worker %d\n", val, ind)
				}
			}
		}(i)
	}
}
