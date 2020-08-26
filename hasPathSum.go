package main

import "fmt"

type TreeNode struct {
	   Val int
	     Left *TreeNode
	     Right *TreeNode
}

func hasPathSum(root *TreeNode, sum int) bool {
	if root == nil {
		return false
	}
	if root.Val == sum && root.Left == nil && root.Right == nil {
		return true
	}
	if root.Left != nil && hasPathSum(root.Left, sum - root.Val) {
		return true
	}
	if root.Right != nil && hasPathSum(root.Right, sum - root.Val) {
		return true
	}
	return false
}

func main() {
	root := &TreeNode{Val:5, Left:nil, Right:nil}
	root.Left = &TreeNode{Val:4}
	root.Right = &TreeNode{Val:8}
	root.Left.Left = &TreeNode{Val:11}
	root.Right.Left = &TreeNode{Val:13}
	root.Right.Right = &TreeNode{Val:4}
	root.Left.Left.Left = &TreeNode{Val:7}
	root.Left.Left.Right = &TreeNode{Val:2}
	root.Right.Right.Right = &TreeNode{Val:1}
	fmt.Println(hasPathSum(root, 22))
}