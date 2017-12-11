package main

import (
	"testing"
	"fmt"
)

// key must be distinct
func TestBuildAvlBST(t *testing.T) {
	data := [][]int{
		{5, 6, 8, 3, 2, 4, 7},
	}
	for _, item := range data {
		avl := BuildAvlBST(item)
		if !avl.Check() {
			fmt.Println(avl)
			t.Errorf("BuildAvlBST(%v) violate avl property.\n", item)
		}
	}
}
