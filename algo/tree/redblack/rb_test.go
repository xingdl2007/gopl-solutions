package main

import (
	"testing"
	"fmt"
	"sort"
	"time"
	"math/rand"
)

func TestBuildRBTree(t *testing.T) {
	data := [][]int{
		{5, 6, 8, 3, 2, 4, 7},
		{1, 2, 3, 4, 5, 6, 7, 8, 9, 10},
	}
	for _, item := range data {
		rb := BuildRBTree(item)
		if !rb.Check() {
			fmt.Println(rb)
			t.Errorf("BuildAvlBST(%v) violate red black property.\n", item)
		}
	}
}

func TestTreeNode_DeleteMin(t *testing.T) {
	data := []struct {
		values []int
		key    int
	}{
		{[]int{5, 6, 8, 3, 2, 4, 7}, -1},
		{[]int{8, 4, 10, 3, 6, 9, 11}, -1},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9,}, -1},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8}, -1},
		{[]int{0, -1, -2, -3, -4, -5, -6, -7, -8}, -1},
		{[]int{}, -1},
		{[]int{1}, -1},
		{[]int{1, 2}, -1},
		{[]int{2, 1}, -1},
		{[]int{5, 3, 6}, -1},
		{[]int{5, 3, 6}, -1},
		{[]int{44, 17, 62, 50, 78, 48, 54, 88}, -1},
	}
	for _, item := range data {
		for cnt := 0; cnt < 100; cnt++ {
			tree := BuildRBTree(item.values)
			cnt := len(item.values)
			sort.Ints(item.values)

			for i := 0; i < cnt; i++ {
				item.key = item.values[0] // min()
				tree = DeleteMin(tree)
				if tree.Check() {
					// make sure key is deleted and all others is still in the tree
					for _, v := range item.values {
						if v != item.key {
							if !tree.Exists(v) {
								t.Errorf("Delete(%d) of %v, %d is missing\n", item.key, item.values, v)
							}
						} else if tree.Exists(item.key) {
							t.Errorf("Delete(%d) of %v, %d is still in avl\n", item.key, item.values, item.key)
						}
					}
				} else {
					t.Errorf("Delete(%d) of %v, violate red black property\n", item.key, item.values)
				}
				item.values = item.values[1:]
			}
		}
	}
}

func TestTreeNode_DeleteMax(t *testing.T) {
	data := []struct {
		values []int
		key    int
	}{
		{[]int{5, 6, 8, 3, 2, 4, 7}, -1},
		{[]int{8, 4, 10, 3, 6, 9, 11}, -1},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9,}, -1},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8}, -1},
		{[]int{0, -1, -2, -3, -4, -5, -6, -7, -8}, -1},
		{[]int{}, -1},
		{[]int{1}, -1},
		{[]int{1, 2}, -1},
		{[]int{2, 1}, -1},
		{[]int{5, 3, 6}, -1},
		{[]int{5, 3, 6}, -1},
		{[]int{44, 17, 62, 50, 78, 48, 54, 88}, -1},
	}
	for _, item := range data {
		for cnt := 0; cnt < 100; cnt++ {
			tree := BuildRBTree(item.values)
			cnt := len(item.values)
			sort.Sort(sort.Reverse(sort.IntSlice(item.values)))

			for i := 0; i < cnt; i++ {
				item.key = item.values[0] // min()
				tree = DeleteMax(tree)
				if tree.Check() {
					// make sure key is deleted and all others is still in the tree
					for _, v := range item.values {
						if v != item.key {
							if !tree.Exists(v) {
								t.Errorf("Delete(%d) of %v, %d is missing\n", item.key, item.values, v)
							}
						} else if tree.Exists(item.key) {
							t.Errorf("Delete(%d) of %v, %d is still in avl\n", item.key, item.values, item.key)
						}
					}
				} else {
					t.Errorf("Delete(%d) of %v, violate red black property\n", item.key, item.values)
				}
				item.values = item.values[1:]
			}
		}
	}
}
func shuffle(data []int) {
	rand.Seed(time.Now().UnixNano())
	for i := len(data) - 1; i > 0; i-- {
		j := rand.Intn(i + 1)
		data[i], data[j] = data[j], data[i]
	}
}

func TestTreeNode_Delete(t *testing.T) {
	data := []struct {
		values []int
		key    int
	}{
		{[]int{5, 6, 8, 3, 2, 4, 7}, -1},
		{[]int{8, 4, 10, 3, 6, 9, 11}, -1},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9,}, -1},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8}, -1},
		{[]int{0, -1, -2, -3, -4, -5, -6, -7, -8}, -1},
		{[]int{}, -1},
		{[]int{1}, -1},
		{[]int{1, 2}, -1},
		{[]int{2, 1}, -1},
		{[]int{5, 3, 6}, -1},
		{[]int{5, 3, 6}, -1},
		{[]int{44, 17, 62, 50, 78, 48, 54, 88}, -1},
	}
	for _, item := range data {
		for cnt := 0; cnt < 100; cnt++ {
			tree := BuildRBTree(item.values)
			cnt := len(item.values)
			shuffle(item.values)

			for i := 0; i < cnt; i++ {
				item.key = item.values[0]
				tree = Delete(tree, item.key)
				if tree.Check() {
					// make sure key is deleted and all others is still in the tree
					for _, v := range item.values {
						if v != item.key {
							if !tree.Exists(v) {
								t.Errorf("Delete(%d) of %v, %d is missing\n", item.key, item.values, v)
							}
						} else if tree.Exists(item.key) {
							t.Errorf("Delete(%d) of %v, %d is still in avl\n", item.key, item.values, item.key)
						}
					}
				} else {
					t.Errorf("Delete(%d) of %v, violate red black property\n", item.key, item.values)
				}
				item.values = item.values[1:]
			}
		}
	}
}
