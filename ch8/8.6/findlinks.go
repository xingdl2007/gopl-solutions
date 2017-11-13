// Copyright © 2017 xingdl2007@gmail.com
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 243.

// Crawl3 crawls web links starting with the command-line arguments.
//
// This version uses bounded parallelism.
// For simplicity, it does not address the termination problem.
//
package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"gopl.io/ch5/links"
	"strings"
)

var depth = flag.Int("depth", 1,
	"Only URLs reachable by at most `depth` links will be fetched")

func crawl(url string) []string {
	fmt.Println(url)
	list, err := links.Extract(url)
	if err != nil {
		log.Print(err)
	}
	return list
}

//!+
func main() {
	flag.Parse()
	worklist := make(chan []string)  // lists of URLs, may have duplicates
	unseenLinks := make(chan string) // de-duplicated URLs
	var n, d int                     // number of pending sends to worklist

	n++
	// record items per depth
	counter := make([]int, *depth+2)
	counter[d] = n

	// Add command-line arguments to worklist.
	// skip depth parameter
	go func() {
		if strings.HasPrefix(os.Args[1], "http") {
			worklist <- os.Args[1:]
		} else {
			worklist <- os.Args[2:]
		}
	}()

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		go func() {
			for link := range unseenLinks {
				foundLinks := crawl(link)
				go func() { worklist <- foundLinks }()
			}
		}()
	}

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist

		// drain worklist then close unseenLinks
		if d > *depth {
			continue
		}
		for _, link := range list {
			if !seen[link] {
				n++ // counter++
				counter[d+1]++

				seen[link] = true
				unseenLinks <- link
			}
		}
		if counter[d]--; counter[d] == 0 {
			d++
		}
	}
	close(unseenLinks)
}

//!-
