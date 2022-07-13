package main

import "fmt"

func main() {
	ch1 := make(chan int) // init fst channel
	ch2 := make(chan int) // init snd channel
	a := [10]int{}        // init array to range over

	// sending integers to ch1
	go func() {
		for i := range a {
			ch1 <- i
		}
		close(ch1) // close ch1
	}()

	// sending integers to ch2, doubling them
	go func() {
		for val := range ch1 {
			ch2 <- val * 2
		}
		close(ch2) // close ch2
	}()

	// print out ch2 values
	for val := range ch2 {
		fmt.Println(val)
	}
}
