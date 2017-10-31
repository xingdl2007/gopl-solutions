// Copyright © 2017 xingdl2007@gmail.com
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 77.

// Netflag demonstrates an integer type used as a bit field.
package main

import "fmt"

const (
	_   = 1 << (10 * iota)
	KiB  // 1024
	MiB
	GiB
	TiB  // overflow 32 int
	PiB
	EiB
	ZiB  // overflow 64 int
	YiB
)

// seems the only valid way
const (
	KB = 1000
	MB = 1000 * KB
	GB = 1000 * MB
	TB = 1000 * GB
	PB = 1000 * TB
	EB = 1000 * PB
	ZB = 1000 * EB
	YB = 1000 * ZB
)

func main() {
	fmt.Printf("%T %[1]v\n", KiB)
	fmt.Printf("%T %[1]v\n", MiB)
	fmt.Printf("%T %[1]v\n", GiB)
	fmt.Printf("%T %[1]v\n", TiB)
	fmt.Printf("%T %[1]v\n", PiB)
	fmt.Printf("%T %[1]v\n", EiB)
	//fmt.Printf("%T %[1]v\n", ZiB)

	fmt.Printf("%T %[1]v\n", KB)
	fmt.Printf("%T %[1]v\n", MB)
	fmt.Printf("%T %[1]v\n", GB)
	fmt.Printf("%T %[1]v\n", TB)
	fmt.Printf("%T %[1]v\n", PB)
	fmt.Printf("%T %[1]v\n", EB)
	//fmt.Printf("%T %[1]v\n", ZB)

	// impressive
	fmt.Println(YiB / ZiB)
	fmt.Println(YB / ZB)

	//
	var f float64 = 1 + 0i
	f = 'a'
	fmt.Println(f)
}

//!-
