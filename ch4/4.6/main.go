// Copyright © 2017 xingdl2007@gmail.com
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 96.

// Dedup prints only one instance of each line; duplicates are removed.
package main

import (
	"unicode/utf8"
	"unicode"
	"fmt"
)

// squashSpace squashes each rune of adjacent Unicode spaces in
// UTF-8 encoded []byte slice into a single ASCII space
func squashSpace(bytes []byte) []byte {
	out := bytes[:0]
	var last rune

	for i := 0; i < len(bytes); {
		r, s := utf8.DecodeRune(bytes[i:])

		if !unicode.IsSpace(r) {
			out = append(out, bytes[i:i+s]...)
		} else if unicode.IsSpace(r) && !unicode.IsSpace(last) {
			// ASCII space literal, untyped constant, will transfer
			// to byte when assign to append's parameter
			// ' ' is int32/rune type, character in Go is actually
			// Unicode code point, but string will utf8 encoded
			out = append(out, ' ')
		}
		last = r
		i += s
	}
	return out
}

//!+
func main() {
	a := 'a'
	r := rune(' ')
	fmt.Printf("%T %v %T %v\n", a, a, r, r)

	// string is []byte
	s := "这个\n割裂的\n世界->   hel\n\v\f\tlo \n世界\t\v wo rld \f! \n\v!\n !"
	// output show utf8 encoded
	fmt.Println(len(s), len([]rune(s)))

	b := squashSpace([]byte(s))
	fmt.Println(string(b))
}

//!-
