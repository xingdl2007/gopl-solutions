// Copyright © 2017 xingdl2007@gmail.com
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 86.

// Rev reverses a slice.
package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	//fmt.Println([]byte("世"))
	//fmt.Println([]byte("界"))

	//!+slice
	s := []byte("hello 世界 world!")
	reverse(s)
	fmt.Println(string(s)) // "!dlrow 界世 olleh"
	//!-slice

	s = []byte("hello 世界 world!")
	reverse2(s)
	fmt.Println(string(s))
}

//!+rev
// reverse reverses a []byte slice (UTF8-encoded string) in place.
// version 1:
func reverse(in []byte) {
	buf := make([]byte, 0, len(in))
	i := len(in)

	for i > 0 {
		_, s := utf8.DecodeLastRune(in[:i])
		buf = append(buf, in[i-s:i]...)
		i -= s
	}
	copy(in, buf)
}

// version 2:
// Can you do it without allocating memory?
// Based on UTF8's property: prefix code
// though this method is not general, should have more elegant solution
func reverse2(in []byte) {
	// first treat as non utf8-encoded data
	for i, j := 0, len(in)-1; i < j; i, j = i+1, j-1 {
		in[i], in[j] = in[j], in[i]
	}

	// try to decode according to utf8, then fix error
	i := 0
	for i < len(in) {
		var tryTwo, tryThree, tryFour bool
		for {
			r, s := utf8.DecodeRune(in[i:])
			if r != utf8.RuneError {
				i += s
				break
			} else {
				// try two byte length, swap two bytes
				if !tryTwo {
					tryTwo = true
					in[i], in[i+1] = in[i+1], in[i]
					continue
				}

				// try three byte length, swap three bytes
				if !tryThree {
					// cancel tryTwo side effect
					in[i], in[i+1] = in[i+1], in[i]
					tryThree = true
					in[i], in[i+2] = in[i+2], in[i]
					continue
				}

				// try four byte length, swap four bytes
				if !tryFour {
					// cancel tryThree side effect
					in[i], in[i+1], in[i+2] = in[i+2], in[i+1], in[i]

					tryFour = true
					in[i], in[i+1], in[i+2], in[i+3] = in[i+3], in[i+2], in[i+1], in[i]
					continue
				}

				// should not be here
				panic("Should not be here!")
			}
		}
	}
}

//!-rev
