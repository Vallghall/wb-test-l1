package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	ch := make(chan int) // init channel
	a := [10000]int{}    // init array to range over
	n := 1               // number of seconds in which the program will stop executing

	// timeout shutdown
	time.AfterFunc(time.Duration(n)*time.Second, func() {
		fmt.Println("time out")
		os.Exit(0)
	})

	// func for sequentially sending data to channel
	go func() {
		for i := range a {
			ch <- i
		}
		close(ch) // closing the channel to range over it
	}()

	// reading all values from channel
	for val := range ch {
		fmt.Printf("value from chan: %d\n", val)
	}
}
