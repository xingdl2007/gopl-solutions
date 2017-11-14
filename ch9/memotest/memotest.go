// Copyright © 2017 xingdl2007@gmail.com
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 272.

// Package memotest provides common functions for
// testing various designs of the memo package.
package memotest

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"testing"
	"time"
)

//!+httpRequestBody
func httpGetBody(url string, cancel <-chan struct{}) (interface{}, error) {
	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}

//!-httpRequestBody

var HTTPGetBody = httpGetBody

func incomingURLs() <-chan string {
	ch := make(chan string)
	go func() {
		for _, url := range []string{
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"http://gopl.io",
			"https://golang.org",
			"https://godoc.org",
			"https://play.golang.org",
			"http://gopl.io",
		} {
			ch <- url
		}
		close(ch)
	}()
	return ch
}

type M interface {
	Get(key string, cancel <-chan struct{}) (interface{}, error)
}

/*
//!+seq
	m := memo.New(httpGetBody)
//!-seq
*/

func Sequential(t *testing.T, m M) {
	//!+seq
	for url := range incomingURLs() {
		start := time.Now()
		value, err := m.Get(url, nil)
		if err != nil {
			log.Print(err)
			continue
		}
		fmt.Printf("%s, %s, %d bytes\n",
			url, time.Since(start), len(value.([]byte)))
	}
	//!-seq
}

/*
//!+conc
	m := memo.New(httpGetBody)
//!-conc
*/

func Concurrent(t *testing.T, m M) {
	//!+conc
	var n sync.WaitGroup
	for url := range incomingURLs() {
		n.Add(1)
		go func(url string) {
			defer n.Done()
			start := time.Now()
			value, err := m.Get(url, nil)
			if err != nil {
				log.Print(err)
				return
			}
			fmt.Printf("%s, %s, %d bytes\n",
				url, time.Since(start), len(value.([]byte)))
		}(url)
	}
	n.Wait()
	//!-conc
}

//output: some is canceled
//
//=== RUN   TestConcurrentCancel
//http://gopl.io, 1.392470116s, 4154 bytes
//http://gopl.io, 1.393739377s, 4154 bytes
//2017/11/14 13:33:26 memo: https://play.golang.org is canceled
//2017/11/14 13:33:26 memo: https://play.golang.org is canceled
//2017/11/14 13:33:26 memo: https://golang.org is canceled
//2017/11/14 13:33:26 memo: https://golang.org is canceled
//2017/11/14 13:33:26 memo: https://godoc.org is canceled
//2017/11/14 13:33:26 memo: https://godoc.org is canceled
//--- PASS: TestConcurrentCancel (1.93s)
//PASS
//ok      gopl.io/gopl-solutions/ch9/9.3  2.998s
func ConcurrentCancel(t *testing.T, m M) {
	//!+conc
	var n sync.WaitGroup

	cancel := make(chan struct{})
	mom := time.NewTimer(1600 * time.Millisecond) // 1.6s
	go func() {
		<-mom.C
		close(cancel)
		mom.Stop()
	}()

	for url := range incomingURLs() {
		n.Add(1)
		go func(url string) {
			defer n.Done()
			start := time.Now()
			value, err := m.Get(url, cancel)
			if err != nil {
				log.Print(err)
				return
			}
			fmt.Printf("%s, %s, %d bytes\n",
				url, time.Since(start), len(value.([]byte)))
		}(url)
	}
	n.Wait()
	//!-conc
}
