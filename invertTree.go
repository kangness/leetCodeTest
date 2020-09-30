package main


type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
}

func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	var right, left *TreeNode
	if root.Left != nil {
		left = invertTree(root.Left)
	}
	if root.Right != nil {
		right = invertTree(root.Right)
	}
	root.Right = left
	root.Left = right
	return root
}

func main () {
	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val:2}
	root.Right = &TreeNode{Val:7}
	root.Left.Left = &TreeNode{Val:1}
	root.Left.Right = &TreeNode{Val:3}
	root.Right.Left = &TreeNode{Val:6}
	root.Right.Right = &TreeNode{Val:9}
	invertTree(root)
}