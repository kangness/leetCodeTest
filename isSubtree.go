package main

import "fmt"

type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
}

func isSameTree(s, t *TreeNode) bool {
	if s == nil && t == nil {
		fmt.Println("dididdididdi")
		return true
	}
	if s == nil || t == nil {
		return false
	}
	if s.Val != t.Val {
		return false
	}
	if !isSameTree(s.Left, t.Left) {
		return false
	}
	if !isSameTree(s.Right, t.Right) {
		return false
	}
	fmt.Println("dididdididdi")
	return true
}
func isSubTree(s, t *TreeNode) bool {
	if s == nil && t == nil {
		return true
	}
	if s == nil && t != nil {
		return false
	}
	if isSameTree(s, t) {
		fmt.Println("----iiiii")
		return true
	}
	if isSubTree(s.Left, t) {
		fmt.Println("====iiiii")
		return true
	}
	if isSubTree(s.Right, t) {
		fmt.Println("====999999")
		return true
		return true
	}
	return false
}

func main () {
	/*
	root := &TreeNode{Val:3}
	root.Left = &TreeNode{Val:4}
	root.Right = &TreeNode{Val:5}
	root.Left.Left = &TreeNode{Val:1}
	root.Left.Right = &TreeNode{Val:2}
	 */
	root := &TreeNode{Val:1}
	root.Left = &TreeNode{Val:2}
	root.Right = &TreeNode{Val:3}
	t := &TreeNode{Val:1}
	t.Left = &TreeNode{Val:2}
	fmt.Println(isSubTree(root, t))
}
