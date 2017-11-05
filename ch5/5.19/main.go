// Copyright © 2017 xingdl2007@gmail.com
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 148.

// Fetch saves the contents of a URL into a local file.
package main

import "fmt"

//!-
func dummy() (ret int, err error) {
	defer func() {
		p := recover()
		ret = 1
		err = fmt.Errorf("internal error: %v", p)
	}()

	panic("panic with no reason, just panic!")
}

func main() {
	fmt.Println(dummy())
}
