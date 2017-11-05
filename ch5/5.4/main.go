// Copyright © 2017 xingdl2007@gmail.com
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 122.
//!+main

// Findlinks1 prints the links in an HTML document read from standard input.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

var mapping = map[string]string{"a": "href", "img": "src", "script": "src", "link": "href"}

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}

	// anchor link
	for _, link := range visit("a", nil, doc) {
		fmt.Println(link)
	}

	fmt.Println("--------------------------------------------------")
	// image
	for _, link := range visit("img", nil, doc) {
		fmt.Println(link)
	}

	fmt.Println("--------------------------------------------------")
	// script
	for _, link := range visit("script", nil, doc) {
		fmt.Println(link)
	}

	fmt.Println("--------------------------------------------------")
	// stylesheet
	for _, link := range visit("link", nil, doc) {
		fmt.Println(link)
	}
}

//!-main

//!+visit
// visit appends to links each link found in n and returns the result.
func visit(target string, links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == target {
		for _, a := range n.Attr {
			if a.Key == mapping[target] {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(target, links, c)
	}
	return links
}

//!-visit

/*
//!+html
package html

type Node struct {
	Type                    NodeType
	Data                    string
	Attr                    []Attribute
	FirstChild, NextSibling *Node
}

type NodeType int32

const (
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
)

type Attribute struct {
	Key, Val string
}

func Parse(r io.Reader) (*Node, error)
//!-html
*/
