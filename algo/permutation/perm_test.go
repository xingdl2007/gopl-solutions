package main

import "testing"

var nums = []int{1, 2, 3, 4, 5,}

func BenchmarkRecursive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Recursive(nums)
	}
}

// faster
func BenchmarkBacktracking(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Backtracking(nums)
	}
}

// faster
func BenchmarkHeap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Heap(nums)
	}
}

func BenchmarkLexicographic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Lexicographic(nums)
	}
}

func BenchmarkJohnsonTrotter(b *testing.B) {
	for i := 0; i < b.N; i++ {
		JohnsonTrotter(5)
	}
}
