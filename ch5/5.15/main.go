// Copyright © 2017 xingdl2007@gmail.com
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 142.

// The sum program demonstrates a variadic function.
package main

import "fmt"

//!+
func sum(vals ...int) int {
	total := 0
	for _, val := range vals {
		total += val
	}
	return total
}

func max(v int, vals ...int) (m int) {
	m = v
	for _, val := range vals {
		if m < val {
			m = val
		}
	}
	return
}

func min(v int, vals ...int) (m int) {
	m = v
	for _, val := range vals {
		if m > val {
			m = val
		}
	}
	return
}

//!-

func main() {
	//!+main
	fmt.Println(sum())           //  "0"
	fmt.Println(sum(3))          //  "3"
	fmt.Println(sum(1, 2, 3, 4)) //  "10"
	//!-main

	//!+slice
	values := []int{1, 2, 3, 4}
	fmt.Println(sum(values...)) // "10"
	//!-slice

	fmt.Println(max(1, 2, 3))
	fmt.Println(min(1, 2, 3))
}
