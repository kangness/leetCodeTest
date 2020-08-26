package main

import "fmt"

type TreeNode struct {
	Val int
	Right *TreeNode
	Left *TreeNode
}

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return cmp(root.Right,root.Left)
}

func cmp(right, left *TreeNode) bool {
	if right == nil && left == nil {
		return true
	}
	if right == nil || left == nil {
		return false
	}
	if right.Val != left.Val {
		return false
	}
	if cmp(right.Left, left.Right) && cmp(right.Right, left.Left) {
		return true
	}
	return false
}

func main() {
	root := &TreeNode{Val:1}
	root.Left = &TreeNode{Val:2}
	root.Right = &TreeNode{Val:2}
	root.Left.Left = &TreeNode{Val:3}
	root.Left.Right = &TreeNode{Val:4}
	root.Right.Left = &TreeNode{Val:4}
	root.Right.Right = &TreeNode{Val:3}
	fmt.Println(isSymmetric(root))
}