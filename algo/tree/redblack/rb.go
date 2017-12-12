package main

// red black tree (2-3 tree) implementation, reference: Algorithms 4th, Robert Sedgewick
import (
	"fmt"
	"bytes"
	"strconv"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
	Color bool
}

const (
	RED   = true
	BLACK = false
)

func (t *TreeNode) isRed() bool {
	if t == nil {
		return false
	}
	return t.Color
}

func rotateLeft(t *TreeNode) *TreeNode {
	x := t.Right
	t.Right = x.Left
	x.Left = t
	x.Color = t.Color
	t.Color = RED
	return x
}

func rotateRight(t *TreeNode) *TreeNode {
	x := t.Left
	t.Left = x.Right
	x.Right = t
	x.Color = t.Color
	t.Color = RED
	return x
}

func flipColors(t *TreeNode) {
	t.Color = !t.Color
	t.Left.Color = !t.Left.Color
	t.Right.Color = !t.Right.Color
}

func balance(t *TreeNode) *TreeNode {
	if t.Right.isRed() && !t.Left.isRed() {
		t = rotateLeft(t)
	}
	if t.Left.isRed() && t.Left.Left.isRed() {
		t = rotateRight(t)
	}
	if t.Left.isRed() && t.Right.isRed() {
		flipColors(t)
	}
	return t
}

func Insert(root *TreeNode, key int) *TreeNode {
	root = insert(root, key)
	// force root be black at the end of insertion
	root.Color = BLACK
	return root
}

func insert(t *TreeNode, key int) *TreeNode {
	if t == nil {
		return &TreeNode{key, nil, nil, RED}
	}
	if key < t.Val {
		t.Left = insert(t.Left, key)
	} else {
		t.Right = insert(t.Right, key)
	}
	return balance(t)
}

// delete is tricky
func DeleteMin(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if !root.Left.isRed() && !root.Right.isRed() {
		root.Color = RED
	}
	if root = deleteMin(root); root != nil {
		root.Color = BLACK
	}
	return root
}

func deleteMin(t *TreeNode) *TreeNode {
	if t.Left == nil {
		return nil
	}
	// borrow nodes from parent or sibling
	if !t.Left.isRed() && !t.Left.Left.isRed() {
		t = moveRedLeft(t)
	}
	t.Left = deleteMin(t.Left)
	return balance(t)
}

func moveRedLeft(t *TreeNode) *TreeNode {
	flipColors(t)
	if t.Right.Left.isRed() {
		t.Right = rotateRight(t.Right)
		t = rotateLeft(t)
		flipColors(t)
	}
	return t
}

func DeleteMax(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if !root.Left.isRed() && !root.Right.isRed() {
		root.Color = RED
	}
	if root = deleteMax(root); root != nil {
		root.Color = BLACK
	}
	return root
}

func deleteMax(t *TreeNode) *TreeNode {
	if t.Left.isRed() {
		t = rotateRight(t)
	}

	if t.Right == nil {
		return nil
	}
	// borrow nodes from parent or sibling
	if !t.Right.isRed() && !t.Right.Left.isRed() {
		t = moveRedRight(t)
	}
	t.Right = deleteMax(t.Right)
	return balance(t)
}

func moveRedRight(t *TreeNode) *TreeNode {
	flipColors(t)
	if t.Left.Left.isRed() {
		t = rotateRight(t)
		flipColors(t)
	}
	return t
}

func min(t *TreeNode) *TreeNode {
	if t == nil {
		return nil
	}
	for t.Left != nil {
		t = t.Left
	}
	return t
}

func Delete(root *TreeNode, key int) *TreeNode {
	if root == nil {
		return nil
	}
	if !root.Left.isRed() && !root.Right.isRed() {
		root.Color = RED
	}
	if root = delete(root, key); root != nil {
		root.Color = BLACK
	}
	return root
}

// difficult
func delete(root *TreeNode, key int) *TreeNode {
	if key < root.Val {
		if !root.Left.isRed() && !root.Left.Left.isRed() {
			root = moveRedLeft(root)
		}
		root.Left = delete(root.Left, key)
	} else {
		if root.Left.isRed() {
			root = rotateRight(root)
		}
		// found
		if key == root.Val && root.Right == nil {
			return nil
		}
		if !root.Right.isRed() && !root.Right.Left.isRed() {
			root = moveRedRight(root)
		}
		if key == root.Val {
			x := min(root.Right)
			root.Val = x.Val
			root.Right = deleteMin(root.Right)
		} else {
			root.Right = delete(root.Right, key)
		}
	}
	return balance(root)
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

func (t *TreeNode) isBST(min, max int) bool {
	if t == nil {
		return true
	}
	if t.Val <= min || t.Val >= max {
		return false
	}
	return t.Left.isBST(min, t.Val) && t.Right.isBST(t.Val, max)
}

// check red black property
func (t *TreeNode) isRBTree() bool {
	// root must be black
	if t.isRed() {
		return false
	}
	var height int
	tmp := t
	for tmp != nil {
		if !tmp.isRed() {
			height++
		}
		tmp = tmp.Left
	}
	return t.redEdge() && t.blackEdge(height)
}

// 1. all left-red edge, no right-red edge
// 2. one node can't have two red edges
func (t *TreeNode) redEdge() bool {
	if t == nil {
		return true
	}
	if t.Right.isRed() {
		return false
	}
	if t.isRed() && t.Left.isRed() {
		return false
	}
	return t.Left.redEdge() && t.Right.redEdge()
}

// black balanced ?
func (t *TreeNode) blackEdge(height int) bool {
	if t == nil {
		return height == 0
	}
	if !t.isRed() {
		height--
	}
	return t.Left.blackEdge(height) && t.Right.blackEdge(height)
}

func (t *TreeNode) Check() bool {
	return t.isBST(math.MinInt64, math.MaxInt64) && t.isRBTree()
}

func BuildRBTree(data []int) *TreeNode {
	var root *TreeNode
	for _, item := range data {
		root = Insert(root, item)
	}
	return root
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
	if t.isRed() {
		buf.WriteString(prefix + addition1 + strconv.Itoa(t.Val) + "(r)\n")
	} else {
		buf.WriteString(prefix + addition1 + strconv.Itoa(t.Val) + "\n")
	}

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
	//data := []int{-1, -2, -3, -4, -5, -6, -7, -8}
	data := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	tree := BuildRBTree(data)

	fmt.Println(tree)
	fmt.Println(tree.Check())

	tree = Delete(tree, 2)
	fmt.Println(tree)
	fmt.Println(tree.Check())
}
