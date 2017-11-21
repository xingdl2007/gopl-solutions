// Copyright © 2017 xingdl2007@gmail.com
// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 165.

// Package intset provides a set of integers based on a bit vector.
package intset

import (
	"bytes"
	"fmt"

	"gopl.io/ch2/popcount"
)

const size = 32 << (^uint(0) >> 63)

//!+intset

// An IntSet is a set of small non-negative integers.
// Its zero value represents the empty set.
type IntSet struct {
	words []uint
	count int
}

// helper
func NewIntSet() *IntSet {
	return new(IntSet)
}

// Has reports whether the set contains the non-negative value x.
func (s *IntSet) Has(x int) bool {
	word, bit := x/size, uint(x%size)
	return word < len(s.words) && s.words[word]&(1<<bit) != 0
}

// Add adds the non-negative value x to the set.
func (s *IntSet) Add(x int) {
	word, bit := x/size, uint(x%size)
	for word >= len(s.words) {
		s.words = append(s.words, 0)
	}
	s.count++
	s.words[word] |= 1 << bit
}

func (s *IntSet) AddAll(x ...int) {
	for _, i := range x {
		s.Add(i)
	}
}

// UnionWith sets s to the union of s and t.
func (s *IntSet) UnionWith(t *IntSet) {
	for i, tword := range t.words {
		if i < len(s.words) {
			s.words[i] |= tword
		} else {
			s.count += popcount.PopCount(uint64(tword))
			s.words = append(s.words, tword)
		}
	}
}

func (s *IntSet) Elems() []int {
	if s.Len() == 0 {
		return nil
	}
	r := make([]int, 0, s.Len())
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < size; j++ {
			if word&(1<<uint(j)) != 0 {
				r = append(r, size*i+j)
			}
		}
	}
	return r
}

// IntersectWith return intersection of s and t.
func (s *IntSet) IntersectWith(t *IntSet) *IntSet {
	r := NewIntSet()
	for _, i := range t.Elems() {
		if s.Has(i) {
			r.Add(i)
		}
	}
	return r
}

// DifferenceWIth return difference of s and t.
func (s *IntSet) DifferenceWith(t *IntSet) *IntSet {
	r := NewIntSet()
	for _, i := range s.Elems() {
		if !t.Has(i) {
			r.Add(i)
		}
	}
	return r
}

// SymmetricDifference return symmetric difference of s and t.
func (s *IntSet) SymmtericDifference(t *IntSet) *IntSet {
	u := s.Copy()
	u.UnionWith(t)

	r := u.DifferenceWith(s.IntersectWith(t))

	fmt.Println(u, s.IntersectWith(t), r)
	return r
}

// return the number of elements
func (s *IntSet) Len() int {
	return s.count
}

// remove x from the set
func (s *IntSet) Remove(x int) {
	if s.Has(x) {
		word, bit := x/size, uint(x%size)
		s.words[word] &^= 1 << bit
		s.count--
	}
}

// remove all elements from the set
func (s *IntSet) Clear() {
	s.words = nil
	s.count = 0
}

// return a copy of the set
func (s *IntSet) Copy() *IntSet {
	c := NewIntSet()
	c.words = make([]uint, len(s.words))
	copy(c.words, s.words)
	c.count = s.count
	return c
}

//!-intset

//!+string

// String returns the set as a string of the form "{1 2 3}".
func (s *IntSet) String() string {
	var buf bytes.Buffer
	buf.WriteByte('{')
	for i, word := range s.words {
		if word == 0 {
			continue
		}
		for j := 0; j < size; j++ {
			if word&(1<<uint(j)) != 0 {
				if buf.Len() > len("{") {
					buf.WriteByte(' ')
				}
				fmt.Fprintf(&buf, "%d", size*i+j)
			}
		}
	}
	buf.WriteByte('}')

	return buf.String()
}

//!-string
