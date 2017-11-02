// Copyright © 2017 xingdl2007@gmail.com
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 97.
//!+

// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int) // counts of Unicode characters
	input := bufio.NewScanner(os.Stdin)

	// set SplitFunc with ScanWords instead of default ScanLines
	// ScanWords just return space-separated words
	input.Split(bufio.ScanWords)

	for input.Scan() {
		counts[input.Text()]++
	}

	if err := input.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "WordCount: %v", err)
		os.Exit(1)
	}

	for k, v := range counts {
		fmt.Println(k, v)
	}
}

//!-
