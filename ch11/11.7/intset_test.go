// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

package intset

import (
	"testing"
	"math/rand"
	"math"
	"time"
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

func generateRandomData(rng *rand.Rand, size int) []int {
	data := make([]int, 0, size)
	for i := 0; i < size; i++ {
		data = append(data, rng.Intn(math.MaxInt16))
	}
	return data
}

var temp []int
var seed = time.Now().UTC().UnixNano()
var rng = rand.New(rand.NewSource(seed))

func BenchmarkIntSet_Add(b *testing.B) {
	if len(temp) == 0 {
		temp = generateRandomData(rng, 1000)
	}
	var x = NewIntSet()
	for i := 0; i < b.N; i++ {
		x.AddAll(temp...)
	}
}

func BenchmarkHashSet_Add(b *testing.B) {
	if len(temp) == 0 {
		temp = generateRandomData(rng, 1000)
	}
	var x = NewHashSet()
	for i := 0; i < b.N; i++ {
		x.AddAll(temp...)
	}
}

func BenchmarkIntSet_UnionWith(b *testing.B) {
	var s1, s2 IntSet
	data := generateRandomData(rng, 1000)
	data2 := generateRandomData(rng, 1000)
	s1.AddAll(data...)
	s2.AddAll(data2...)

	for i := 0; i < b.N; i++ {
		s1.UnionWith(&s2)
	}
}

func BenchmarkHashSet_UnionWith(b *testing.B) {
	var s1, s2 = NewHashSet(), NewHashSet()
	data := generateRandomData(rng, 1000)
	data2 := generateRandomData(rng, 1000)
	s1.AddAll(data...)
	s2.AddAll(data2...)

	for i := 0; i < b.N; i++ {
		s1.UnionWith(s2)
	}
}
