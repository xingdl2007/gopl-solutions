// Copyright © 2017 xingdl2007@gmail.com
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package main

import (
	"time"
	"fmt"
)

func main() {
	ping, pong := make(chan int), make(chan int)
	var counter int64

	t := time.NewTimer(1 * time.Minute)
	done := make(chan struct{})
	shutdown := make(chan struct{})

	// ping->pong
	go func() {
	loop:
		for {
			select {
			case <-shutdown:
				break loop
			case v := <-ping:
				counter++
				pong <- v
			}
		}
		done <- struct{}{}
	}()

	go func() {
	loop:
		for {
			select {
			case <-shutdown:
				break loop
			case v := <-pong:
				ping <- v
			}
		}
		done <- struct{}{}
	}()

	// kick it start
	ping <- 1

	// 1 minutes
	<-t.C
	close(shutdown)

	// drain the one which is blocked
	select {
	case <-ping:
	case <-pong:
	}

	<-done
	<-done
	t.Stop()
	fmt.Println(counter)
}
