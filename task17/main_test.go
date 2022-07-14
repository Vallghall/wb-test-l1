package main

import "testing"

func BenchmarkBinarySearchRecursive(b *testing.B) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := 0; i < b.N; i++ {
		binarySearchRecursive(a, 2)
	}
}

func BenchmarkBinarySearchIterative(b *testing.B) {
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := 0; i < b.N; i++ {
		binarySearchIterative(a, 2)
	}
}
