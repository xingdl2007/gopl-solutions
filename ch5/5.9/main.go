// Copyright © 2017 xingdl2007@gmail.com
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package main

import (
	"strings"
	"fmt"
)

// expand replace "$foo" in s with foo("foo") and return the result
func expand(s string, f func(string) string) string {
	items := strings.Split(s, " ")
	for i, item := range items {
		if strings.HasPrefix(item, "$") {
			items[i] = f(item[1:])
		}
	}
	return strings.Join(items, " ")
}

func main() {
	s := "$Hello $world"
	f := func(s string) string {
		return s + s
	}

	fmt.Println(expand(s, f))
}
