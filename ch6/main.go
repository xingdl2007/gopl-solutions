package main

import (
	"gopl.io/gopl-solutions/ch6/intset"
	"fmt"
)

func main() {
	s := intset.NewIntSet()
	s.Clear()

	s.AddAll(1, 2, 3, 4, 5, 6)
	c := s.Copy()
	fmt.Println(c.Len())

	t := intset.NewIntSet()
	t.AddAll(7, 8, 9, 1, 2, 3)

	fmt.Println(s, t)

	// union test
	c.UnionWith(t)
	fmt.Println(c)

	// intersect
	fmt.Println(s.IntersectWith(t))

	fmt.Println(s, s.Len())
	// difference
	fmt.Println(s.DifferenceWith(t))
	// symmetric difference
	fmt.Println(s.SymmtericDifference(t))

	// remove test
	fmt.Println(c)
	c.Remove(1)
	c.Remove(1)
	c.Remove(9)
	c.Remove(1900)
	fmt.Println(c)
	c.Clear()
	fmt.Println(c)
}
