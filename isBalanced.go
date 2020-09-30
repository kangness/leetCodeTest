package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}

func isBalance(root *TreeNode) bool {
	if root == nil {
		return true
	}
	lh := getHight(root.Left)
	rh := getHight(root.Right)
	if math.Abs(float64(lh) -float64(rh)) > 1 {
		return false
	}
	if root.Left != nil {
		if isBalance(root.Left) == false {
			return false
		}
	}
	if root.Right != nil {
		if isBalance(root.Right) == false {
			return false
		}
	}
	return true
}

func getHight(root *TreeNode) int {
	if root == nil {
		return 0
	}
	var lh,rh int
	if root.Left != nil {
		lh = getHight(root.Left)
	}
	if root.Right != nil {
		rh = getHight(root.Right)
	}
	if lh > rh {
		return lh + 1
	}else {
		return rh + 1
	}
}

func main() {
	/*
	root := &TreeNode{Val:3}
	root.Left = &TreeNode{Val:9}
	root.Right = &TreeNode{Val:20}
	root.Right.Left = &TreeNode{Val:15}
	root.Right.Right = &TreeNode{Val:7}
	 */
	fmt.Println(getHight(root))
	fmt.Println(isBalance(root))
}
