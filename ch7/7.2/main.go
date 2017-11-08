// Copyright © 2017 xingdl2007@gmail.com
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 173.

package main

import (
	"io"
	"os"
	"fmt"
)

type CounterWriter struct {
	counter int64
	writer  io.Writer
}

// must be pointer type in order to count
func (cw *CounterWriter) Write(p []byte) (int, error) {
	cw.counter += int64(len(p))
	return cw.writer.Write(p)
}

// newWriter is a Writer Wrapper, return original Writer
// and a Counter which record bytes have written
func CountingWriter(w io.Writer) (io.Writer, *int64) {
	cw := CounterWriter{0, w}
	return &cw, &cw.counter
}

// CounterWriter is a proxy
func main() {
	cw, c := CountingWriter(os.Stdout)
	fmt.Fprintf(cw, "This is %dth day without a job.\n", 26)
	fmt.Println(*c)

	// another output
	fmt.Fprint(cw, "Fuck...\n")
	fmt.Println(*c)
}
