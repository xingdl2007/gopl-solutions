// Copyright © 2017 xingdl2007@gmail.com
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package main

import (
	"net"
	"log"
	"io"
	"os"
	"strings"
)

func mustCopy(w io.Writer, r io.Reader) {
	if _, err := io.Copy(w, r); err != nil {
		log.Fatal(err)
	}
}

func getTime(addr string) {
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	mustCopy(os.Stdout, conn)
}

// parameter: NewYork=localhost:8010 London=localhost:8020 Tokyo=localhost:8030
func main() {
	for _, locate := range os.Args[2:] {
		s := strings.Split(locate, "=")
		go getTime(s[1])
	}

	// main goroutine
	s := strings.Split(os.Args[1], "=")
	getTime(s[1])
}
