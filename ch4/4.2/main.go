// Copyright © 2017 xingdl2007@gmail.com
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 83.

// The sha256 command computes the SHA256 hash (an array) of a string.
package main

//!+
import (
	"bufio"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
	"strconv"
)

var len = flag.String("n", "256", "Generate sha256/384/512 digest")

func main() {
	flag.Parse()
	algo, err := strconv.Atoi(*len)

	// parameter validation
	if err != nil ||
		algo != 256 && algo != 384 && algo != 512 {
		algo = 256
	}

	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		text := input.Text()
		switch algo {
		case 256:
			fmt.Printf("%x\n", sha256.Sum256([]byte(text)))
		case 384:
			fmt.Printf("%x\n", sha512.Sum384([]byte(text)))
		case 512:
			fmt.Printf("%x\n", sha512.Sum512([]byte(text)))
		default:
			fmt.Printf("%x\n", sha256.Sum256([]byte(text)))
		}

		// ignore Err
	}
}

//!-
