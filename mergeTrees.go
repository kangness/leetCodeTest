package main
import ("fmt")

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func mergeTrees(t1, t2 *TreeNode) *TreeNode {
	if t1 == nil {
		return t2
	}
	if t2 == nil {
		return t1
	}
	t := &TreeNode{}
	if t1 != nil {
		t.Val = t.Val + t1.Val
	}
	if t2 != nil {
		t.Val = t.Val + t2.Val
	}
	t.Right = mergeTrees(t1.Right, t2.Right)
	t.Left = mergeTrees(t1.Left, t2.Left)
	return t
}


