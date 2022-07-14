package main

import "fmt"

func main() {
	x, y := 5, 10     // init two values
	fmt.Println(x, y) // print them out
	x, y = y, x       // change places without temporal value
	fmt.Println(x, y) // print out again
}
