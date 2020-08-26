package main

import "fmt"

type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
}
var treeMap map[int]*TreeNode
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	right := minDepth(root.Right)
	left := minDepth(root.Left)
	if right < left {
		return right + 1
	}else {
		return left + 1
	}
}
func main() {
	root := &TreeNode{}
	root.Val = 3
	root.Left = &TreeNode{Val:9,Left:nil,Right:nil}
	root.Right = &TreeNode{Val:20}
	root.Right.Left = &TreeNode{Val:15,Left:nil,Right:nil}
	root.Right.Right = &TreeNode{Val:7,Left:nil,Right:nil}
	fmt.Println(minDepth(root))
}
