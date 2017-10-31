// Copyright © 2017 xingdl2007@gmail.com
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
// 	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
// 	1
// 	12
// 	123
// 	1,234
// 	1,234,567,890
//
package main

import (
	"fmt"
	"os"
	"bytes"
	"strings"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

//!+
// comma inserts commas in a non-negative decimal integer string.
//func comma(s string) string {
//	n := len(s)
//	if n <= 3 {
//		return s
//	}
//	return comma(s[:n-3]) + "," + s[n-3:]
//}

// comma inserts commas in a non-negative decimal integer string.
func comma(s string) string {
	var sign string

	// skip optional sign flag
	if hasSign := strings.HasPrefix(s, "+") ||
		strings.HasPrefix(s, "-"); hasSign {
		sign = s[0:1]
		s = s[1:]
	}

	var buf bytes.Buffer
	buf.WriteString(sign)

	if period := strings.LastIndex(s, "."); period != -1 {
		doWork(s[:period], &buf)
		buf.WriteByte('.')
		doWork(s[period+1:], &buf)
	} else {
		doWork(s, &buf)
	}

	return buf.String()
}

// add to only one buf
func doWork(s string, buf *bytes.Buffer) {
	n := len(s)
	if n <= 3 {
		buf.WriteString(s)
		return
	}

	d, t := len(s)%3, len(s)/3
	if d != 0 {
		buf.WriteString(s[:d])
		buf.WriteByte(',')
	}

	for i := 0; i < t; i++ {
		buf.WriteString(s[i*3+d:d+i*3+3])
		if i != t-1 {
			buf.WriteByte(',')
		}
	}
}

//!-
