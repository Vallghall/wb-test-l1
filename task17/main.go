package main

import "fmt"

func main() {
	a := []int{10, 23, 35, 46, 57, 68, 72, 81, 94}
	r := binarySearchRecursive(a, 57)
	i := binarySearchIterative(a, 57)
	fmt.Printf("recursive: %d\niterative: %d\n", r, i)
}

// recursive implementation
func binarySearchRecursive[T ~int](a []T, val T) int {
	mid := len(a) / 2 // middle index

	// check equality
	if a[mid] == val {
		return mid
	}

	// recursive call on the right half
	if a[mid] < val {
		return binarySearchRecursive(a[mid:], val)
	}

	// recursive call on the left half
	if a[mid] > val {
		return binarySearchRecursive(a[:mid], val)
	}

	// if value does not exist in array
	return -1
}

func binarySearchIterative[T ~int](a []T, val T) int {
	low, high := 0, len(a)-1 // indexes for leftmost and rightmost values
	mid := high / 2          // middle element index

	// start searching
	for low <= high {
		// check for equality
		if a[mid] == val {
			return mid
		}

		// change left and mid to search in the right half
		if a[mid] < val {
			low = mid + 1
			mid = (high - low) / 2
			continue
		}

		// change mid and right to search in the left half
		high = mid - 1
		mid = (high - low) / 2
	}

	// if value does not exist in array
	return -1
}
