// Copyright © 2017 xingdl2007@gmail.com
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 86.

// Rev reverses a slice.
package main

import (
	"fmt"
)

func main() {
	//!+array
	a := [...]int{0, 1, 2, 3, 4, 5}
	fmt.Println(a) // "[5 4 3 2 1 0]"
	//!-array

	// rotateLeft
	rotateLeft(a[:], 2)
	fmt.Println(a) // "[2 3 4 5 0 1]"

	// rotateRight
	rotateRight(a[:], 2)
	fmt.Println(a) // "[0 1 2 3 4 5]"

	//!+slice
	s := []int{0, 1, 2, 3, 4, 5}
	// Rotate s left by two positions.
	reverse(s[:2])
	reverse(s[2:])
	reverse(s)
	fmt.Println(s) // "[2 3 4 5 0 1]"
	//!-slice
}

//!+rev
// reverse reverses a slice of ints in place.
func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

// rotateLeft rotate slice s left
func rotateLeft(s []int, n int) []int {
	n = n % len(s)
	if n > 0 && n < len(s) {
		temp := make([]int, n)
		copy(temp, s[:n])

		copy(s, s[n:])
		copy(s[len(s)-n:], temp)
	}
	return s
}

func rotateRight(s []int, n int) []int {
	n = n % len(s)
	if n > 0 && n < len(s) {
		temp := make([]int, n)
		copy(temp, s[len(s)-n:])

		copy(s[n:], s)
		copy(s, temp)
		return s
	}
	return s
}

//!-rev
