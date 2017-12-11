package main

import (
	"testing"
	"sort"
	"math/rand"
	"time"
)

func TestBuildBinarySearchTree(t *testing.T) {
	data := [][]int{
		{8, 4, 10, 3, 6, 9, 11},
	}
	for _, item := range data {
		tree := BuildBinarySearchTree(item)
		if !tree.Check() {
			t.Errorf("build bst from %v failed, violate bst proterty\n", tree)
		}
	}
}

func TestTreeNode_Min(t *testing.T) {
	data := []struct {
		values   []int
		expected int
	}{
		{[]int{8, 4, 10, 3, 6, 9, 11}, 3,},
	}
	for _, item := range data {
		tree := BuildBinarySearchTree(item.values)
		if tree.Min() != item.expected {
			t.Errorf("Min(%v) is %d, expected %d \n", item.values, tree.Min(), item.expected)
		}
	}
}

func TestTreeNode_Max(t *testing.T) {
	data := []struct {
		values   []int
		expected int
	}{
		{[]int{8, 4, 10, 3, 6, 9, 11}, 11,},
	}
	for _, item := range data {
		tree := BuildBinarySearchTree(item.values)
		if tree.Max() != item.expected {
			t.Errorf("Max(%v) is %d, expected %d \n", item.values, tree.Max(), item.expected)
		}
	}
}

func TestTreeNode_Floor(t *testing.T) {
	data := []struct {
		values   []int
		key      int
		expected interface{}
	}{
		{[]int{8, 4, 10, 3, 6, 9, 11}, 2, nil},
		{[]int{8, 4, 10, 3, 6, 9, 11}, 5, 4},
		{[]int{8, 4, 10, 3, 6, 9, 11}, 15, 11},
	}
	for _, item := range data {
		tree := BuildBinarySearchTree(item.values)
		ret := tree.Floor(item.key)

		// nil == nil is legal, operator == is not defined on nil
		if ret == nil && item.expected != nil || ret != nil && item.expected == nil {
			t.Errorf("Floor(%v) is %d, expected %d \n", item.values, ret, item.expected)
		} else if ret != nil && item.expected != nil && ret.val != item.expected {
			t.Errorf("Floor(%v) is %d, expected %d \n", item.values, ret.val, item.expected)
		}
	}
}

func TestTreeNode_Ceiling(t *testing.T) {
	data := []struct {
		values   []int
		key      int
		expected interface{}
	}{
		{[]int{8, 4, 10, 3, 6, 9, 11}, 12, nil},
		{[]int{8, 4, 10, 3, 6, 9, 11}, 5, 6},
		{[]int{8, 4, 10, 3, 6, 9, 11}, 7, 8},
	}
	for _, item := range data {
		tree := BuildBinarySearchTree(item.values)
		ret := tree.Ceiling(item.key)

		// nil == nil is legal, operator == is not defined on nil
		if ret == nil && item.expected != nil || ret != nil && item.expected == nil {
			t.Errorf("Ceiling(%v) is %d, expected %d \n", item.values, ret, item.expected)
		} else if ret != nil && item.expected != nil && ret.val != item.expected {
			t.Errorf("Ceiling(%v) is %d, expected %d \n", item.values, ret.val, item.expected)
		}
	}
}

func TestTreeNode_DeleteMin(t *testing.T) {
	data := []struct {
		values []int
		min    int
	}{
		{[]int{8, 4, 10, 3, 6, 9, 11}, 3},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9,}, 0},
	}
	for _, item := range data {
		tree := BuildBinarySearchTree(item.values)
		sort.Ints(item.values)
		cnt := len(item.values)
		for i := 0; i < cnt; i++ {
			item.min = tree.Min()
			tree = tree.DeleteMin()
			if tree.Check() {
				// make sure min is deleted and all others is still in
				for _, v := range item.values {
					if v != item.min {
						if !tree.Exists(v) {
							t.Errorf("DeleteMin(%d) of %v, %d is missing\n", item.min, item.values, v)
						}
					} else if tree.Exists(item.min) {
						t.Errorf("DeleteMin(%d) of %v, %d is still in bst\n", item.min, item.values, item.min)
					}
				}
			} else {
				t.Errorf("DeleteMin(%d) of %v, violate bst property\n", item.min, item.values)
			}
			item.values = item.values[1:]
		}
	}
}

func TestTreeNode_DeleteMax(t *testing.T) {
	data := []struct {
		values []int
		max    int
	}{
		{[]int{8, 4, 10, 3, 6, 9, 11}, 11},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9,}, 9},
	}
	for _, item := range data {
		tree := BuildBinarySearchTree(item.values)
		// reverse order
		sort.Sort(sort.Reverse(sort.IntSlice(item.values)))
		cnt := len(item.values)
		for i := 0; i < cnt; i++ {
			item.max = tree.Max()
			tree = tree.DeleteMax()

			if tree.Check() {
				// make sure max is deleted and all others is still in
				for _, v := range item.values {
					if v != item.max {
						if !tree.Exists(v) {
							t.Errorf("DeleteMax(%d) of %v, %d is missing\n", item.max, item.values, v)
						}
					} else if tree.Exists(item.max) {
						t.Errorf("DeleteMax(%d) of %v, %d is still in bst\n", item.max, item.values, item.max)
					}
				}
			} else {
				t.Errorf("DeleteMax(%d) of %v, violate bst property\n", item.max, item.values)
			}
			item.values = item.values[1:]
		}
	}
}

// very useful
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
		{[]int{8, 4, 10, 3, 6, 9, 11}, -1},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9,}, -1},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8}, -1},
		{[]int{0, -1, -2, -3, -4, -5, -6, -7, -8}, -1},
		{[]int{}, -1},
		{[]int{1}, -1},
		{[]int{1, 2}, -1},
		{[]int{2, 1}, -1},
		{[]int{5, 3, 6}, -1},
	}
	for _, item := range data {
		tree := BuildBinarySearchTree(item.values)
		cnt := len(item.values)
		shuffle(item.values)

		for i := 0; i < cnt; i++ {
			item.key = item.values[0]
			tree = tree.Delete(item.key)
			if tree.Check() {
				// make sure max is deleted and all others is still in
				for _, v := range item.values {
					if v != item.key {
						if !tree.Exists(v) {
							t.Errorf("Delete(%d) of %v, %d is missing\n", item.key, item.values, v)
						}
					} else if tree.Exists(item.key) {
						t.Errorf("Delete(%d) of %v, %d is still in bst\n", item.key, item.values, item.key)
					}
				}
			} else {
				t.Errorf("Delete(%d) of %v, violate bst property\n", item.key, item.values)
			}
			item.values = item.values[1:]
		}
	}
}
