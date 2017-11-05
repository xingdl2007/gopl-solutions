// Copyright © 2017 xingdl2007@gmail.com
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package main

import (
	"strings"
	"fmt"
)

// call strings.Join?
func VariadicJoin(sep string, a ...string) string {
	return strings.Join(a, sep)
}

func VariadicJoin2(sep string, a ...string) string {
	switch len(a) {
	case 0:
		return ""
	case 1:
		return a[0]
	case 2:
		return a[0] + sep + a[1]
	case 3:
		return a[0] + sep + a[1] + sep + a[2]
	}

	n := len(sep) * (len(a) - 1)
	for i := 0; i < len(a); i++ {
		n += len(a[i])
	}

	b := make([]byte, n)
	bp := copy(b, a[0])
	for _, s := range a[1:] {
		bp += copy(b[bp:], sep)
		bp += copy(b[bp:], s)
	}
	return string(b)
}

func main() {
	fmt.Println(VariadicJoin(", ", "hello", "world"))
}
