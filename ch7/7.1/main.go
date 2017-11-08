// Copyright © 2017 xingdl2007@gmail.com
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 173.

// Bytecounter demonstrates an implementation of io.Writer that counts bytes.
package main

import (
	"fmt"
	"bufio"
	"bytes"
	"unicode/utf8"
	"unicode"
)

//!+bytecounter

type WordCounter int

// ScanWords is a split function for a Scanner that returns each
// space/punct-separated word of text, with surrounding spaces deleted. It will
// never return an empty string. The definition of space is set by
// unicode.IsSpace.
func ScanWords(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// Skip leading spaces/punct.
	start := 0
	for width := 0; start < len(data); start += width {
		var r rune
		r, width = utf8.DecodeRune(data[start:])
		if !unicode.IsSpace(r) && !unicode.IsPunct(r) {
			break
		}
	}
	// Scan until space, marking end of word.
	for width, i := 0, start; i < len(data); i += width {
		var r rune
		r, width = utf8.DecodeRune(data[i:])
		if unicode.IsSpace(r) || unicode.IsPunct(r) {
			return i + width, data[start:i], nil
		}
	}
	// If we're at EOF, we have a final, non-empty, non-terminated word. Return it.
	if atEOF && len(data) > start {
		return len(data), data[start:], nil
	}
	// Request more data.
	return start, nil, nil
}

func (c *WordCounter) Write(p []byte) (int, error) {
	// range a string, you get every rune
	s := bufio.NewScanner(bytes.NewBuffer(p))

	// word split
	s.Split(ScanWords)
	for s.Scan() {
		// debug only
		//fmt.Println(s.Text())
		*c++
	}

	return len(p), s.Err()
}

func (c *WordCounter) String() string {
	return fmt.Sprintf("%d word(s)", *c)
}

// line counter

type LineCounter int

func (l *LineCounter) Write(p []byte) (int, error) {
	s := bufio.NewScanner(bytes.NewBuffer(p))

	// line counter
	for s.Scan() {
		*l++
	}
	return len(p), s.Err()
}

func (l *LineCounter) String() string {
	return fmt.Sprintf("%d line(s)", *l)
}

//!-wordcounter && linecounter
func main() {
	//!+main
	var c WordCounter
	c.Write([]byte("hello, 世界! 哈哈,哈"))
	fmt.Println(&c) // "3 words"

	c = 0 // reset the counter
	var name = "Dolly"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(&c) // "2 words"

	var l LineCounter
	l.Write([]byte("你\n好，世\n界\n\n你好\n吗，世界\n"))
	fmt.Println(&l)

	//!-main
}
