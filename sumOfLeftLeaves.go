package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}
func sumOfLeftLeaves(root *TreeNode) int {
	var result int
	if root == nil {
		return result
	}
	if root != nil && root.Left != nil {
		if root.Left.Left == nil && root.Left.Right == nil {
			result = result + root.Left.Val
		}
		result = result + sumOfLeftLeaves(root.Left)
	}
	if root != nil && root.Right != nil {
		result = result + sumOfLeftLeaves(root.Right)
	}
	return result
}


func main() {
	root := &TreeNode{Val:3}
	root.Left = &TreeNode{Val:9}
	root.Right = &TreeNode{Val:20}
	root.Right.Left = &TreeNode{Val:15}
	root.Right.Right = &TreeNode{Val:7}
	fmt.Println(sumOfLeftLeaves(root))
}
