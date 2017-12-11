package main

// build an AVL tree
import (
	"fmt"
	"strings"
	"bytes"
	"strconv"
	"math"
)

// Definition for a binary tree node.
/*
	Note: height is optional, if not valid, the following function
    can be used to calculate height on the fly (recursive):

	func (t *TreeNode) Height() int {
		if t == nil {
			return -1
		}
		return max(t.Left.Height(), t.Right.Height()) + 1
	}

	Then, there will be no need to AdjustHeight() in rotation and insert.
 */
type TreeNode struct {
	Val    int
	Left   *TreeNode
	Right  *TreeNode
	height int
}

// nil is valid receiver for convenience
func (t *TreeNode) BalanceFactor() int {
	if t == nil {
		return 0
	}
	return t.Left.Height() - t.Right.Height()
}

// nil is valid receiver for convenience
func (t *TreeNode) Height() int {
	if t == nil {
		return -1
	}
	return t.height
}

func (t *TreeNode) AdjustHeight() {
	t.height = max(t.Left.Height(), t.Right.Height()) + 1
}

// helper function
func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func Insert(t *TreeNode, val int) *TreeNode {
	if t == nil {
		return &TreeNode{Val: val}
	}
	if val < t.Val {
		t.Left = Insert(t.Left, val)
	} else {
		t.Right = Insert(t.Right, val)
	}
	t.AdjustHeight()

	// reBalance if Insert break AVL invariant; t is the nearest node
	if t.BalanceFactor() > 1 {
		// R or LR
		if t.Left.BalanceFactor() > 0 {
			t = rotateRight(t)
		} else {
			t = rotateLeftRight(t)
		}
	} else if t.BalanceFactor() < -1 {
		// L or RL
		if t.Right.BalanceFactor() < 0 {
			t = rotateLeft(t)
		} else {
			t = rotateRightLeft(t)
		}
	}

	return t
}

func rotateRight(t *TreeNode) *TreeNode {
	c := t.Left
	t.Left = c.Right
	c.Right = t

	t.AdjustHeight()
	c.AdjustHeight()
	return c
}

func rotateLeft(t *TreeNode) *TreeNode {
	c := t.Right
	t.Right = c.Left
	c.Left = t

	t.AdjustHeight()
	c.AdjustHeight()
	return c
}

func rotateLeftRight(r *TreeNode) *TreeNode {
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

func rotateRightLeft(l *TreeNode) *TreeNode {
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

func (t *TreeNode) isBST(min, max int) bool {
	if t == nil {
		return true
	}
	if t.Val <= min || t.Val >= max {
		return false
	}
	return t.Left.isBST(min, t.Val) && t.Right.isBST(t.Val, max)
}

func (t *TreeNode) isAVL() bool {
	if t == nil {
		return true
	}
	if t.BalanceFactor() > 1 || t.BalanceFactor() < -1 {
		return false
	}
	return t.Left.isAVL() && t.Right.isAVL()
}

func (t *TreeNode) Check() bool {
	return t.isBST(math.MinInt64, math.MaxInt64) && t.isAVL()
}

func BuildAvlBST(a []int) *TreeNode {
	var root *TreeNode
	for _, item := range a {
		root = Insert(root, item)
	}
	return root
}

// in order traverse
func (t *TreeNode) Squash() string {
	var nodes []string
	var walk func(*TreeNode)
	walk = func(n *TreeNode) {
		if n == nil {
			return
		}
		walk(n.Left)
		nodes = append(nodes, fmt.Sprintf("%d[%d:%d] ", n.Val, n.Height(), n.BalanceFactor()))
		walk(n.Right)
	}
	nodes = append(nodes, "[")
	walk(t)
	nodes = append(nodes, "]")
	return strings.Join(nodes, " ")
}

// useful and elegant
func (t *TreeNode) Print(prefix string, isTail bool) string {
	if t == nil {
		return ""
	}
	var buf bytes.Buffer
	var addition1, addition2 string
	if isTail {
		addition1 = "└── "
		addition2 = "    "
	} else {
		addition1 = "├── "
		addition2 = "│   "
	}
	buf.WriteString(prefix + addition1 + strconv.Itoa(t.Val) + "\n")

	// left node is tail if there is no right node
	if t.Right != nil {
		isTail = false
	} else {
		isTail = true
	}
	buf.WriteString(t.Left.Print(prefix+addition2, isTail))
	// right node is always the tail
	buf.WriteString(t.Right.Print(prefix+addition2, true))
	return buf.String()
}

func (t *TreeNode) String() string {
	return t.Print("", true)
}

func main() {
	array1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	array2 := []int{9, 8, 7, 6, 5, 4, 3, 2, 1}
	fmt.Println(BuildAvlBST(array1))
	fmt.Println(BuildAvlBST(array2))
}
