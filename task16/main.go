package main

import "fmt"

func main() {
	a := []int{4, 7, 2, 8, 3, 6, 5, 1, 0, 5} // init slice
	fmt.Println(a)
	quicksort(a) // sort slice
	fmt.Println(a)
}

func quicksort[T ~int](a []T) {
	//if slice's len < 2, it doesn't need sorting
	if len(a) < 2 {
		return
	}

	// indexes of the leftmost and rightmost slice elements
	left, right := 0, len(a)-1

	// pivot index
	center := len(a) / 2

	// swap positions of pivot and rightmost element for convenient range over
	a[center], a[right] = a[right], a[center]

	for i, _ := range a {
		// if an element is lesser than a pivot, it goes left
		if a[i] < a[right] {
			// send element on the left half of the slice
			a[i], a[left] = a[left], a[i]
			// inc the left index for swap positions
			left++
		}
	}

	// return pivot to a place between "lesser than pivot"
	// and "bigger than pivot" elements
	a[left], a[right] = a[right], a[left]

	quicksort(a[:left])   // recursively sort left side
	quicksort(a[left+1:]) // recursively sort right side
}
