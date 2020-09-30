package main

import (
	"fmt"
	"math"
	"sync"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

var maxValue = int(math.MinInt64)
func maxPathSum2(root *TreeNode) int {
	sync.Mutex{}
	if root == nil {
		return 0
	}
	left := maxPathSum2(root.Left)
	right := maxPathSum2(root.Right)
	if left < 0 {
		left = 0
	}
	if right < 0 {
		right = 0
	}
	if maxValue < left + right + root.Val {
		maxValue = left + right + root.Val
	}
	if left > right {
		return left + root.Val
	}
	return right + root.Val
}

func maxPathSum(root *TreeNode) int {
	maxPathSum2(root)
	return maxValue
}
func main() {
	root := &TreeNode{Val:1}
	root.Left = &TreeNode{Val:2}
	root.Right = &TreeNode{Val:3}
	fmt.Println(maxPathSum(root))
}