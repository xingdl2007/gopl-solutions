// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package memo_test

import (
	"testing"

	"gopl.io/gopl-solutions/ch9/9.3"
	"gopl.io/gopl-solutions/ch9/memotest"
)

var httpGetBody = memotest.HTTPGetBody

func Test(t *testing.T) {
	m := memo.New(httpGetBody)
	defer m.Close()
	memotest.Sequential(t, m)
}

func TestConcurrent(t *testing.T) {
	m := memo.New(httpGetBody)
	defer m.Close()
	memotest.Concurrent(t, m)
}

func TestConcurrentCancel(t *testing.T) {
	m := memo.New(httpGetBody)
	defer m.Close()
	memotest.ConcurrentCancel(t, m)
}
