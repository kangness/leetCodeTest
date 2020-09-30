package main

import (
	"fmt"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

var pre *TreeNode
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if isValidBST(root.Left) {
		if pre == nil || pre.Val < root.Val {
			fmt.Println(pre)
			pre = root
			return isValidBST(root.Right)
		}
	}
	return false
}


func main() {
	/*
		root := &TreeNode{Val:5}
		root.Left = &TreeNode{Val:4}
		root.Right = &TreeNode{Val:8}
		root.Left.Left = &TreeNode{Val:11}
		root.Left.Left.Left = &TreeNode{Val:7}
		root.Left.Left.Right = &TreeNode{Val:2}
		root.Right.Left = &TreeNode{Val:13}
		root.Right.Right = &TreeNode{Val:4}
		root.Right.Right.Left = &TreeNode{Val:5}
		root.Right.Right.Right = &TreeNode{Val:1}
	*/
	root := &TreeNode{Val: 0}
	/*
	root.Left = &TreeNode{Val: 5}
	root.Right = &TreeNode{Val: 15}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 20}
	 */
	fmt.Println(isValidBST(root))

}
