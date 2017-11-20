// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package intset

import (
	"fmt"
	"testing"
)

func TestIntSet(t *testing.T) {
	var intset = NewIntSet()
	var hashset = NewHashSet()

	intset.Add(1)
	intset.Add(144)
	intset.Add(299)

	hashset.Add(1)
	hashset.Add(144)
	hashset.Add(299)

	if intset.String() != hashset.String() {
		t.Errorf("intset(%q) != hashset(%q)", intset.String(), hashset.String())
	}

	intset2 := NewIntSet()
	intset2.Add(42)

	hashset2 := NewHashSet()
	hashset2.Add(42)

	intset.UnionWith(intset2)
	hashset.UnionWith(hashset2)

	if intset.String() != hashset.String() {
		t.Errorf("intset(%q) != hashset(%q)", intset.String(), hashset.String())
	}
}

func Example_one() {
	//!+main
	var x, y IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	fmt.Println(x.String()) // "{1 9 144}"

	y.Add(9)
	y.Add(42)
	fmt.Println(y.String()) // "{9 42}"

	x.UnionWith(&y)
	fmt.Println(x.String()) // "{1 9 42 144}"

	fmt.Println(x.Has(9), x.Has(123)) // "true false"
	//!-main

	// Output:
	// {1 9 144}
	// {9 42}
	// {1 9 42 144}
	// true false
}

func Example_two() {
	var x IntSet
	x.Add(1)
	x.Add(144)
	x.Add(9)
	x.Add(42)

	//!+note
	fmt.Println(&x)         // "{1 9 42 144}"
	fmt.Println(x.String()) // "{1 9 42 144}"
	fmt.Println(x)          // "{[4398046511618 0 65536]}"
	//!-note

	// Output:
	// {1 9 42 144}
	// {1 9 42 144}
	// {[4398046511618 0 65536]}
}
