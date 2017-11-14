// Copyright © 2017 xingdl2007@gmail.com
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package main

import (
	"time"
	"fmt"
)

const N = 1000000

func main() {
	hello := "hello"
	var mid [N]chan string
	for i := 0; i < N; i++ {
		mid[i] = make(chan string)
	}
	head, tail := mid[0], mid[N-1]

	for i := 0; i < N-1; i++ {
		go func(i int) {
			mid[i+1] <- <-mid[i]
		}(i)
	}

	s := time.Now()
	head <- hello
	<-tail
	fmt.Printf("%d goroutines, %fs\n", N, time.Since(s).Seconds())
}
