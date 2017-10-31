// Copyright © 2017 xingdl2007@gmail.com
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package main

import "fmt"

func main() {
	fmt.Println(anagram("今天是个好天气，啊好天气", "好天气个啊好天气是，天今"))
}

// first construct a map with s1, then check s2 with map
// map is rune:int pair
func anagram(s1, s2 string) bool {
	m := make(map[rune]int)
	for _, r := range s1 {
		m[r]++
	}
	for _, r := range s2 {
		if i, ok := m[r]; ok {
			if i > 1 {
				m[r]--
			} else {
				delete(m, r)
			}
		} else {
			return false
		}
	}
	return len(m) == 0
}
