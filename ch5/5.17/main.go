// Copyright © 2017 xingdl2007@gmail.com
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package main

import (
	"golang.org/x/net/html"
	"os"
	"net/http"
	"fmt"
)

func main() {
	for _, url := range os.Args[1:] {
		outline(url)
	}
}

func outline(url string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	doc, err := html.Parse(resp.Body)
	if err != nil {
		return err
	}

	//!+call
	imgs := ElementsByTagName(doc, "img");
	headings := ElementsByTagName(doc, "h1", "h2", "h3", "h4")
	//!-call

	fmt.Println(len(imgs), len(headings))
	return nil
}

func ElementsByTagName(doc *html.Node, name ...string) (out []*html.Node) {
	pre := func(doc *html.Node) {
		for _, n := range name {
			if doc.Type == html.ElementNode && doc.Data == n && doc.FirstChild != nil {
				out = append(out, doc.FirstChild)
			}
		}
	}
	forEachNode(doc, pre, nil)
	return
}

// Copied from gopl.io/ch5/outline2.
func forEachNode(n *html.Node, pre, post func(n *html.Node)) {
	if pre != nil {
		pre(n)
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		forEachNode(c, pre, post)
	}
	if post != nil {
		post(n)
	}
}
