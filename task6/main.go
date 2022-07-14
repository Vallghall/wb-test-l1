package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{} // init sync.WaitGroup
	wg.Add(3)              // inc WG counter

	// init context with cancelFunc
	ctx, cancelFunc := context.WithCancel(context.Background())
	// launch a goroutine that stops when context closes
	go func() {
		defer wg.Done()
		<-ctx.Done() // stop goroutine
		fmt.Println("goroutine stopped")
	}()
	cancelFunc() // cancel the context

	ch := make(chan struct{}) // init channel
	// launch a goroutine that stops when channel receives a value
	go func() {
		defer wg.Done()
		<-ch // stop goroutine
		fmt.Println("goroutine stopped")
	}()
	ch <- struct{}{} // send value to channel
	// launch a goroutine that stops after receiving channel value in 1 second
	go func() {
		defer wg.Done()
		<-time.After(time.Second * 1)
		fmt.Println("goroutine stopped")
	}()
	wg.Wait() // wait for goroutines to do their job
}
