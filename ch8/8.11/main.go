// Copyright © 2017 xingdl2007@gmail.com
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 17.
//!+

// Fetchall fetches URLs in parallel and reports their times and sizes.
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string)
	cancel := make(chan struct{})

	for i, url := range os.Args[1:] {
		fmt.Println(i)
		go fetch(url, cancel, ch) // start a goroutine
	}

	fmt.Println(<-ch) // receive from channel ch
	close(cancel)

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, cancel <-chan struct{}, ch chan<- string) {
	start := time.Now()
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return
	}
	req.Cancel = cancel
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %7d  %s", secs, nbytes, url)
}

//!-
