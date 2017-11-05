// Copyright © 2017 xingdl2007@gmail.com
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package main

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
	"strings"
	"os"
)

// CountWordsAndImages does an HTTP GET request for the HTML
// document url and returns the number of words and images in it.
func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}
	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	// content in <style> or <script> are ignored
	if n.Type == html.ElementNode {
		if n.Data == "style" || n.Data == "script" {
			return
		} else if n.Data == "img" {
			images++
		}
	} else if n.Type == html.TextNode {
		text := strings.TrimSpace(n.Data)
		for _, line := range strings.Split(text, "\n") {
			if line != "" {
				words += len(strings.Split(line, " "))
				//fmt.Printf("%s %q %d\n", line, strings.Split(line, " "), len(strings.Split(line, " ")))
			}
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		w, i := countWordsAndImages(c)
		words += w
		images += i
	}
	return
}

func main() {
	fmt.Println(len(strings.Split("", " ")))
	for _, url := range os.Args[1:] {
		fmt.Println(CountWordsAndImages(url))
	}
}
