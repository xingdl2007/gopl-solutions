// Copyright © 2017 xingdl2007@gmail.com∂∂∂∂
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 165.

// Package intset provides a set of integers based on a bit vector.

package intset

import (
	"sort"
	"bytes"
	"fmt"
)

type HashSet map[int]bool

func NewHashSet() HashSet {
	return make(map[int]bool)
}

func (s HashSet) Has(x int) bool {
	return s[x]
}

func (s HashSet) Add(x int) {
	s[x] = true
}

func (s HashSet) AddAll(x ...int) {
	for _, a := range x {
		s[a] = true
	}
}

func (s HashSet) UnionWith(t HashSet) {
	for k := range t {
		s[k] = true
	}
}

func (s HashSet) String() string {
	var buf bytes.Buffer
	var items []int
	for k := range s {
		items = append(items, k)
	}
	sort.Ints(items)
	buf.WriteByte('{')
	for i, k := range items {
		fmt.Fprintf(&buf, "%d", k)
		if i != len(items)-1 {
			buf.WriteByte(' ')
		}
	}
	buf.WriteByte('}')
	return buf.String()
}
