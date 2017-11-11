package main

// build an AVL tree
import (
	"fmt"
	"strings"
)

// Definition for a binary tree node.
type TreeNode struct {
	Val         int
	Left        *TreeNode
	Right       *TreeNode
	LeftHeight  int // Left subTree Height = left.height() +1
	RightHeight int
}

// nil is valid receiver for convenience
func (t *TreeNode) BalanceFactor() int {
	if t == nil {
		return 0
	}
	return t.LeftHeight - t.RightHeight
}

// nil is valid receiver for convenience
func (t *TreeNode) Height() int {
	if t == nil {
		return -1
	}
	return max(t.LeftHeight, t.RightHeight)
}

func (t *TreeNode) AdjustHeight() {
	t.LeftHeight = t.Left.Height() + 1
	t.RightHeight = t.Right.Height() + 1
}

// helper function
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// info: how to maintain subtree height
func put(t *TreeNode, val int) *TreeNode {
	if t == nil {
		return &TreeNode{Val: val}
	}
	if val < t.Val {
		t.Left = put(t.Left, val)
		//t.LeftHeight = t.Left.Height() + 1
	} else {
		t.Right = put(t.Right, val)
		//t.RightHeight = t.Right.Height() + 1
	}
	t.AdjustHeight()

	// reBalance if put break AVL invariant
	// t is the nearest node
	if t.BalanceFactor() == 2 || t.BalanceFactor() == -2 {
		if val < t.Val {
			// R or RL
			if t.Left.BalanceFactor() == 1 {
				// R
				t = rRotation(t)
			} else {
				// RL
				t = rlRotation(t)
			}
		} else {
			// L or LR
			if t.Right.BalanceFactor() == -1 {
				// L
				t = lRotation(t)
			} else {
				// LR
				t = lrRotation(t)
			}
		}
	}
	return t
}

func rRotation(t *TreeNode) *TreeNode {
	c := t.Left
	t.Left = c.Right
	c.Right = t

	t.AdjustHeight()
	c.AdjustHeight()
	return c
}

func lRotation(t *TreeNode) *TreeNode {
	c := t.Right
	t.Right = c.Left
	c.Left = t

	t.AdjustHeight()
	c.AdjustHeight()
	return c
}

func rlRotation(r *TreeNode) *TreeNode {
	c := r.Left
	g := c.Right

	c.Right = g.Left
	r.Left = g.Right

	g.Left = c
	g.Right = r

	c.AdjustHeight()
	r.AdjustHeight()
	g.AdjustHeight()
	return g
}

func lrRotation(l *TreeNode) *TreeNode {
	c := l.Right
	g := c.Left

	c.Left = g.Right
	l.Right = g.Left

	g.Right = c
	g.Left = l

	l.AdjustHeight()
	c.AdjustHeight()
	g.AdjustHeight()

	return g
}

func BuildAvlBST(a []int) *TreeNode {
	var root *TreeNode
	for _, item := range a {
		root = put(root, item)
	}
	return root
}

func (t *TreeNode) String() string {
	var nodes []string
	var walk func(*TreeNode)
	walk = func(n *TreeNode) {
		if n == nil {
			return
		}
		walk(n.Left)
		nodes = append(nodes, fmt.Sprintf("%d[%d-%d:%d] ", n.Val,
			n.LeftHeight, n.RightHeight, n.LeftHeight-n.RightHeight))
		walk(n.Right)
	}
	nodes = append(nodes, "[")
	walk(t)
	nodes = append(nodes, "]")
	return strings.Join(nodes, " ")
}

func main() {
	array1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	array2 := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println(BuildAvlBST(array1))
	fmt.Println(BuildAvlBST(array2))

	//var test *TreeNode
	//fmt.Println(test == nil)
	//fmt.Println(test.Height())
}
