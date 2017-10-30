// Copyright © 2017 xingdl2007@gmail.com
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// 1.3 measure the time of different implementation of echo
package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func echo1() string {
	start := time.Now()
	defer func() {
		fmt.Printf("echo1: %v ns\n", time.Since(start).Nanoseconds())
	}()

	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	return s
}

func echo2() string {
	start := time.Now()
	defer func() {
		fmt.Printf("echo2: %v ns\n", time.Since(start).Nanoseconds())
	}()

	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	return s
}

func echo3() string {
	start := time.Now()
	defer func() {
		fmt.Printf("echo3: %v ns\n", time.Since(start).Nanoseconds())
	}()

	return strings.Join(os.Args[1:], " ")
}

func main() {
	echo1()
	echo2()
	echo3()
}

// go run main.go a b c d e f g h i j k l m n o p q r s t u v w x y z
// output:
// echo1: 4743 ns
// echo2: 2495 ns
// echo3: 1097 ns
