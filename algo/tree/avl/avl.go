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

func (t *TreeNode) balance() *TreeNode {
	var b = t
	// reBalance if Insert break AVL invariant; t is the nearest node
	if t.BalanceFactor() > 1 {
		// R or LR
		if t.Left.BalanceFactor() > 0 {
			b = rotateRight(t)
		} else {
			b = rotateLeftRight(t)
		}
	} else if t.BalanceFactor() < -1 {
		// L or RL
		if t.Right.BalanceFactor() < 0 {
			b = rotateLeft(t)
		} else {
			b = rotateRightLeft(t)
		}
	}
	return b
}

func (t *TreeNode) Exists(v int) bool {
	if t == nil {
		return false
	}
	switch {
	case v == t.Val:
		return true
	case v < t.Val:
		return t.Left.Exists(v)
	default:
		return t.Right.Exists(v)
	}
}

func Insert(t *TreeNode, key int) *TreeNode {
	if t == nil {
		return &TreeNode{Val: key}
	}
	if key < t.Val {
		t.Left = Insert(t.Left, key)
	} else {
		t.Right = Insert(t.Right, key)
	}
	t.AdjustHeight()
	return t.balance()
}

func (t *TreeNode) min() *TreeNode {
	if t == nil {
		return nil
	}
	if t.Left == nil {
		return t
	}
	return t.Left.min()
}

func (t *TreeNode) deleteMin() *TreeNode {
	if t == nil {
		return nil
	}
	if t.Left == nil {
		return t.Right
	}
	t.Left = t.Left.deleteMin()
	t.AdjustHeight()
	return t.balance()
}

// interesting, much like classic bst deletion
// ref: http://www.cdn.geeksforgeeks.org/avl-tree-set-2-deletion/
func Delete(t *TreeNode, key int) *TreeNode {
	if t == nil {
		return &TreeNode{Val: key}
	}
	if key < t.Val {
		t.Left = Delete(t.Left, key)
	} else if key > t.Val {
		t.Right = Delete(t.Right, key)
	} else {
		if t.Left == nil {
			return t.Right
		}
		if t.Right == nil {
			return t.Left
		}
		x := t
		t = x.Right.min()
		t.Right = Delete(x.Right, t.Val)
		t.Left = x.Left
	}
	t.AdjustHeight()
	return t.balance()
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
	array1 := []int{5, 6, 8, 3, 2, 4, 7}
	tree := BuildAvlBST(array1)

	tree = Delete(tree, 3)
	fmt.Println(tree)
	fmt.Println(tree.Check())

	data := []int{44, 17, 62, 32, 50, 78, 48, 54, 88}
	tree = BuildAvlBST(data)

	fmt.Println(tree)
	tree = Delete(tree, 32)
	fmt.Println(tree)
	fmt.Println(tree.Check())
}
