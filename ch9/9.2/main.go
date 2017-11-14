// Copyright © 2017 xingdl2007@gmail.como
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 45.

// (Package doc comment intentionally malformed to demonstrate golint.)
//!+
package popcount

import "sync"

// pc[i] is the population count of i.
var table [256]byte

func initTable() {
	for i := range table {
		table[i] = table[i/2] + byte(i&1)
	}
}

var load sync.Once

func pc(i byte) byte {
	load.Do(initTable)
	return table[i]
}

// PopCount returns the population count (number of set bits) of x.
func PopCount(x uint64) int {
	return int(pc(byte(x>>(0*8))) +
		pc(byte(x>>(1*8))) +
		pc(byte(x>>(2*8))) +
		pc(byte(x>>(3*8))) +
		pc(byte(x>>(4*8))) +
		pc(byte(x>>(5*8))) +
		pc(byte(x>>(6*8))) +
		pc(byte(x>>(7*8))))
}

//!-
