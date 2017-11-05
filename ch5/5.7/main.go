// Copyright © 2017 xingdl2007@gmail.com
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 133.

// Outline prints the outline of an HTML document tree.
package main

import (
	"fmt"
	"net/http"
	"os"

	"golang.org/x/net/html"
	"strings"
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
	forEachNode(doc, startElement, endElement)
	//!-call

	return nil
}

//!+forEachNode
// forEachNode calls the functions pre(x) and post(x) for each node
// x in the tree rooted at n. Both functions are optional.
// pre is called before the children are visited (preorder) and
// post is called after (postorder).
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

//!-forEachNode

//!+startend
var depth int

func startElement(n *html.Node) {
	if n.Type == html.ElementNode {
		var attribute string
		for _, attr := range n.Attr {
			attribute += fmt.Sprintf("%s=%q ", attr.Key, attr.Val)
		}

		child := ""
		// <div /> is illegal
		if n.Data == "img" && n.FirstChild == nil {
			child = " /"
		}

		if len(attribute) > 1 {
			attribute = attribute[:len(attribute)-1]
			fmt.Printf("%*s<%s %s%s>\n", depth*2, "", n.Data, attribute, child)
		} else {
			fmt.Printf("%*s<%s%s>\n", depth*2, "", n.Data, child)
		}
		depth++
	} else if n.Type == html.TextNode || n.Type == html.CommentNode {
		if !(n.Parent.Type == html.ElementNode &&
			(n.Parent.Data == "script" || n.Parent.Data == "style")) {
			for _, line := range strings.Split(n.Data, "\n") {
				line = strings.TrimSpace(line)
				if line != "" && line != "\n" {
					fmt.Printf("%*s%s\n", depth*2, "", line)
				}
			}
		}
	}
}

func endElement(n *html.Node) {
	if n.Type == html.ElementNode {
		depth--
		// <div /> is illegal
		if !(n.Data == "img" && n.FirstChild == nil) {
			fmt.Printf("%*s</%s>\n", depth*2, "", n.Data)
		}
	}
}

//!-startend
