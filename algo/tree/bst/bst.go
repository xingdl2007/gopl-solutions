package main

import (
	"fmt"
	"strconv"
	"strings"
	"bytes"
	"math"
)

// Definition for a binary search tree node.
type TreeNode struct {
	val         int
	left, right *TreeNode
}

// t can be nil, then a brand new tree is constructed
func (t *TreeNode) Insert(v int) *TreeNode {
	if t == nil {
		return &TreeNode{v, nil, nil}
	}
	if v < t.val {
		t.left = t.left.Insert(v)
	} else {
		t.right = t.right.Insert(v)
	}
	return t
}

func (t *TreeNode) Exists(v int) bool {
	if t == nil {
		return false
	}
	switch {
	case v == t.val:
		return true
	case v < t.val:
		return t.left.Exists(v)
	default:
		return t.right.Exists(v)
	}
}

func (t *TreeNode) Min() int {
	if t.left == nil {
		return t.val
	}
	return t.left.Min()
}

// return min() node
func (t *TreeNode) min() *TreeNode {
	for t.left != nil {
		t = t.left
	}
	return t
}

func (t *TreeNode) Max() int {
	if t.right == nil {
		return t.val
	}
	return t.right.Max()
}

// range related
func (t *TreeNode) Floor(k int) *TreeNode {
	if t == nil {
		return nil
	}
	if k < t.val {
		return t.left.Floor(k)
	}
	if k == t.val {
		return t
	}
	// t is an candidate
	if n := t.right.Floor(k); n != nil {
		return n
	}
	return t
}

func (t *TreeNode) Ceiling(k int) *TreeNode {
	if t == nil {
		return nil
	}
	if k > t.val {
		return t.right.Ceiling(k)
	}
	if k == t.val {
		return t
	}
	if n := t.left.Ceiling(k); n != nil {
		return n
	}
	return t
}

// iterative
//func (t *TreeNode) DeleteMin() *TreeNode {
//	if t == nil {
//		return nil
//	}
//	parent, child := t, t
//	for child.left != nil {
//		parent = child
//		child = child.left
//	}
//	if parent == child {
//		return t.right
//	}
//	parent.left = child.right
//	return t
//}

//func (t *TreeNode) DeleteMax() *TreeNode {
//	if t == nil {
//		return nil
//	}
//	parent, child := t, t
//	for child.right != nil {
//		parent = child
//		child = child.right
//	}
//	if parent == child {
//		return t.left
//	}
//	parent.right = child.left
//	return t
//}

// recursive
func (t *TreeNode) DeleteMin() *TreeNode {
	if t == nil {
		return nil
	}
	if t.left == nil {
		return t.right
	}
	t.left = t.left.DeleteMin()
	return t
}

func (t *TreeNode) DeleteMax() *TreeNode {
	if t == nil {
		return nil
	}
	if t.right == nil {
		return t.left
	}
	t.right = t.right.DeleteMax()
	return t
}

// deletion is little difficult
func (t *TreeNode) Delete(k int) *TreeNode {
	if t == nil {
		return nil
	}
	if k < t.val {
		t.left = t.left.Delete(k)
	} else if k > t.val {
		t.right = t.right.Delete(k)
	} else {
		if t.left == nil {
			return t.right
		}
		if t.right == nil {
			return t.left
		}
		// more elaborate and understandable
		x := t.right.min()
		x.right = t.right.Delete(x.val)
		x.left = t.left
		return x
	}
	return t
}

/*
pretty print a large tree: "Solution is inspired by the "tree" command in linux."
https://stackoverflow.com/questions/4965335/how-to-print-binary-tree-diagram

public class TreeNode {
    final String name;
    final List<TreeNode> children;

    public TreeNode(String name, List<TreeNode> children) {
        this.name = name;
        this.children = children;
    }
    public void print() {
        print("", true);
    }
    private void print(String prefix, boolean isTail) {
        System.out.println(prefix + (isTail ? "└── " : "├── ") + name);
        for (int i = 0; i < children.size() - 1; i++) {
            children.get(i).print(prefix + (isTail ? "    " : "│   "), false);
        }
        if (children.size() > 0) {
            children.get(children.size() - 1)
                    .print(prefix + (isTail ?"    " : "│   "), true);
        }
    }
}
 */
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
	buf.WriteString(prefix + addition1 + strconv.Itoa(t.val) + "\n")

	// left node is tail if there is no right node
	if t.right != nil {
		isTail = false
	} else {
		isTail = true
	}
	buf.WriteString(t.left.Print(prefix+addition2, isTail))
	// right node is always the tail
	buf.WriteString(t.right.Print(prefix+addition2, true))
	return buf.String()
}

func (t *TreeNode) String() string {
	return t.Print("", true)
}

// binary search tree property validation
func (t *TreeNode) isBST(min, max int) bool {
	if t == nil {
		return true
	}
	if t.val <= min || t.val >= max {
		return false
	}
	return t.left.isBST(min, t.val) && t.right.isBST(t.val, max)
}

func (t *TreeNode) Check() bool {
	return t.isBST(math.MinInt64, math.MaxInt64)
}

// Build a binary search tree from a slice
func BuildBinarySearchTree(data []int) *TreeNode {
	var root *TreeNode
	for _, v := range data {
		root = root.Insert(v)
	}
	return root
}

func main() {
	data := []int{8, 4, 10, 3, 6, 9, 11}
	t := BuildBinarySearchTree(data)

	// simple pretty print test
	fmt.Println(t)
}

// try with level print, but the result if not very well
// @Deprecated
func (t *TreeNode) String2() string {
	var p = []*TreeNode{t}
	var q []*TreeNode
	var res [][]string
	for {
		var line = make([]string, 0, len(p))
		var cnt int
		for _, n := range p {
			if n != nil {
				line = append(line, strconv.Itoa(n.val), "  ")
				q = append(q, n.left, n.right)
			} else {
				cnt++
				q = append(q, nil, nil)
				line = append(line, "  ")
			}
		}
		if cnt == len(p) {
			break
		}
		p, q = q, p
		q = q[:0]
		res = append(res, line)
	}

	var str = make([]string, len(res))
	for i := len(res) - 1; i >= 0; i-- {
		str[i] += Space(len(res)-i) + strings.Join(res[i], "")
	}
	return strings.Join(str, "\n")
}

// n * " "
func Space(n int) string {
	if n <= 0 {
		return ""
	}
	var s = make([]string, n)
	ind := "  "
	for i := 0; i < n; i++ {
		s = append(s, ind)
	}
	return strings.Join(s, "")
}
