// Copyright © 2017 xingdl2007@gmail.com
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 278.

// Package memo provides a concurrency-safe non-blocking memoization
// of a function.  Requests for different keys proceed in parallel.
// Concurrent requests for the same key block until the first completes.
// This implementation uses a monitor goroutine.
package memo

import (
	"fmt"
)

//!+Func

// Func is the type of the function to memoize.
type Func func(key string, cancel <-chan struct{}) (interface{}, error)

// A result is the result of calling a Func.
type result struct {
	value interface{}
	err   error
}

type entry struct {
	res   result
	ready chan struct{} // closed when res is ready
}

//!-Func

//!+get

// A request is a message requesting that the Func be applied to key.
type request struct {
	key      string
	cancel   <-chan struct{}
	response chan<- result // the client wants a single result
}

type Memo struct {
	requests chan request
	delete   chan string // buffered
}

// New returns a memoization of f.  Clients must subsequently call Close.
func New(f Func) *Memo {
	memo := &Memo{requests: make(chan request), delete: make(chan string, 10)}
	go memo.server(f)
	return memo
}

func (memo *Memo) Get(key string, cancel <-chan struct{}) (interface{}, error) {
	response := make(chan result)
	memo.requests <- request{key, cancel, response}
	res := <-response
	return res.value, res.err
}

func (memo *Memo) Delete(key string) {
	memo.delete <- key
}

func (memo *Memo) Close() {
	close(memo.delete)
	close(memo.requests)
}

//!-get

//!+monitor

func (memo *Memo) server(f Func) {
	cache := make(map[string]*entry)
loop:
	for {
		select {
		case key, ok := <-memo.delete:
			if !ok {
				break loop
			}
			delete(cache, key)
		case req, ok := <-memo.requests:
			if !ok {
				break loop
			}
			e := cache[req.key]
			if e == nil {
				// This is the first request for this key.
				e = &entry{ready: make(chan struct{})}
				cache[req.key] = e
				go e.call(f, req.key, req.cancel) // call f(key)
			}
			go e.deliver(req.response, req.cancel, memo.delete, req.key)
		}
	}
}

func (e *entry) call(f Func, key string, cancel <-chan struct{}) {
	// Evaluate the function.
	e.res.value, e.res.err = f(key, cancel)

	// Broadcast the ready condition.
	close(e.ready)
}

func (e *entry) deliver(response chan<- result, cancel <-chan struct{},
	delete chan<- string, key string) {
	// Wait for the ready condition.
	<-e.ready

	// if this request is canceled which must be happened before e.ready
	// then delete current entry (the results of a canceled Func call
	// should not be cached)
	select {
	case <-cancel:
		delete <- key
		response <- result{nil, fmt.Errorf("%s is canceled", key)}
	default:
		// Send the result to the client.
		response <- e.res
	}
}

//!-monitor
