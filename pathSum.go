package main

import "fmt"

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}


func pathSum(root *TreeNode, sum int) [][]int {
	if root == nil {
		return nil
	}
	if root.Val == sum && root.Left == nil && root.Right == nil {
		return [][]int{{root.Val}}
	}
	left := pathSum(root.Left, sum - root.Val)
	right := pathSum(root.Right, sum - root.Val)
	fmt.Println(left, right)
	var result [][]int
	if left != nil {
		for _, l := range left {
			var tmp []int
			tmp = append(tmp, root.Val)
			tmp = append(tmp, l...)
			result = append(result, tmp)
		}
	}
	if right != nil {
		for _, l := range right {
			var tmp []int
			tmp = append(tmp, root.Val)
			tmp = append(tmp, l...)
			result = append(result, tmp)
		}
	}
	return result
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
	root := &TreeNode{Val: -2}
	root.Right = &TreeNode{Val: -3}
	fmt.Println(pathSum(root,-5))

}
