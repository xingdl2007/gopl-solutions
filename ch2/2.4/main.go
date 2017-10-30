// Copyright © 2017 xingdl2007@gmail.com
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 45.

// (Package doc comment intentionally malformed to demonstrate golint.)
//!+
package main

import (
	"os"
	"strconv"
	"fmt"
	"log"
)

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) (ret int) {
	for i := 0; i < 64; i++ {
		if x&1 == 1 {
			ret++
		}
		x >>= 1
	}
	return ret
}

func main() {
	for _, num := range os.Args[1:] {
		x, err := strconv.ParseUint(num, 10, 64)
		if err != nil {
			log.Printf("popcount: %v\n", err)
			os.Exit(1)
		}
		fmt.Println(PopCount(x))
	}
}

//!-
