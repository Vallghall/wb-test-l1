package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5, 6}  // init slice
	i := 2                        // index for element to remove
	a = append(a[:i], a[i+1:]...) // removal
	fmt.Println(a)                // print out
}
