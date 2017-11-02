// Copyright © 2017 xingdl2007@gmail.com
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 96.

// Dedup prints only one instance of each line; duplicates are removed.
package main

import (
	"fmt"
)

//!+
func main() {
	fmt.Println([]string{"", "c"})

	// more test
	fmt.Printf("\ndup1: ")
	fmt.Println(filteDup([]string{}))
	fmt.Println(filteDup([]string(nil)))
	fmt.Println(filteDup([]string{"", "c"}))
	fmt.Println(filteDup([]string{"a", "", ""}))
	fmt.Println(filteDup([]string{"hello", "hello", "hello", "world", "世界"}))

	fmt.Printf("\ndup2: \n")
	fmt.Println(filteDup2([]string{}))
	fmt.Println(filteDup2([]string(nil)))
	fmt.Println(filteDup2([]string{"", "c"}))
	fmt.Println(filteDup2([]string{"a", "", ""}))
	fmt.Println(filteDup2([]string{"hello", "hello", "hello", "world", "世界"}))

	fmt.Printf("\ndup3: \n")
	fmt.Println(filteDup3([]string{}))
	fmt.Println(filteDup3([]string(nil)))
	fmt.Println(filteDup3([]string{"", "c"}))
	fmt.Println(filteDup3([]string{"a", "", ""}))
	fmt.Println(filteDup3([]string{"hello", "hello", "hello", "world", "世界"}))
}

// filterDup filter adjacent strings in place
// slice with index
func filteDup(strings []string) []string {
	i := 0
	var last string

	// in case of strings[0] == ""
	if len(strings) > 0 && strings[0] == "" {
		// different with strings[0] is enough
		last = "last"
	}

	for _, s := range strings {
		if s != last {
			strings[i] = s
			last = s
			i++
		}
	}
	return strings[:i]
}

// filterDup filter adjacent strings in place
func filteDup3(s []string) []string {
	if len(s) == 0 {
		return s
	}

	i := 0
	for j := 1; j < len(s); j++ {
		if s[j] != s[i] {
			i++
			s[i] = s[j]
		}
	}
	return s[:i+1]
}

// append() version
func filteDup2(strings []string) []string {
	var last string
	ret := strings[:0]

	// in case of strings[0] == ""
	if len(strings) > 0 && strings[0] == "" {
		// different with strings[0] is enough
		last = "last"
	}

	for _, s := range strings {
		if s != last {
			ret = append(ret, s)
			last = s
		}
	}
	return ret
}

//!-
