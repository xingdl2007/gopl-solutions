// Copyright © 2017 xingdl2007@gmail.com
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package main

import (
	"sort"
	"fmt"
)

// IsPalindrome report whether s is a palindrome
func IsPalindrome(s sort.Interface) bool {
	i, j := 0, s.Len()-1
	for j > i {
		// Less() only
		if !s.Less(i, j) && !s.Less(j, i) {
			i++
			j--
		} else {
			return false
		}
	}
	return true
}

func main() {
	ints := []int{1, 2, 3, 2, 1}
	fmt.Println(IsPalindrome(sort.IntSlice(ints)))

	strings := []string{"hello", "world", "world", "你好"}
	fmt.Println(IsPalindrome(sort.StringSlice(strings)))
}
