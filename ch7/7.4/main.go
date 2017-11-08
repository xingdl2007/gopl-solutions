// Copyright © 2017 xingdl2007@gmail.com
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
package main

import (
	"golang.org/x/net/html"
	"fmt"
	"os"
	"net/http"
	"log"
	"bytes"
	"io"
)

// a simple version of strings.Reader implement io.Reader interface
type StringReader struct {
	s string
	i int64
}

func (r *StringReader) Read(p []byte) (n int, err error) {
	// if p is nil or empty, return (0, nil)
	if len(p) == 0 {
		return 0, nil
	}

	// copy() guarantee copy min(len(p),len(r.s[r.i:])) bytes
	n = copy(p, r.s[r.i:])
	if r.i += int64(n); r.i >= int64(len(r.s)) {
		err = io.EOF
	}
	return
}

// NewReader return a StringReader with s
func NewReader(s string) *StringReader {
	return &StringReader{s, 0}
}

func main() {
	rsp, err := http.Get(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	//var sbuf = make([]byte, 4096)
	//n, err := sr.Read(sbuf)
	//for n > 0 {  	                  // always process n bytes before check io.EOF condition
	//	fmt.Printf("%s", sbuf[:n])
	//	if err != nil {
	//		break
	//	}
	//	n, err = sr.Read(sbuf)
	//}

	// construct string from a reader, use bytes.Buffer as container
	var buf bytes.Buffer
	io.Copy(&buf, rsp.Body)
	defer rsp.Body.Close()

	sr := NewReader(buf.String()) // or strings.NewReader(buf.String())

	doc, err := html.Parse(sr)
	if err != nil {
		fmt.Fprintf(os.Stderr, "outline: %v\n", err)
		os.Exit(1)
	}
	outline(nil, doc)
}

// from ch5/outline
func outline(stack []string, n *html.Node) {
	if n.Type == html.ElementNode {
		stack = append(stack, n.Data) // push tag
		fmt.Println(stack)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		outline(stack, c)
	}
}
