package main

import (
	"container/list"
	"fmt"
)

type TreeNode struct {
	Val int
	Right *TreeNode
	Left *TreeNode
}

func levelOrderBottom(root *TreeNode)[][]int {
	if root == nil {
		return [][]int{}
	}
	var result [][]int
	dataList := list.New()
	tmp := list.New()
	dataList.PushFront(root)
	for {
		var t []int
		tmp.Init()
		for e := dataList.Front(); e != nil; e = e.Next() {
			value := e.Value.(*TreeNode)
			t = append(t, value.Val)
			if value.Left != nil {
				tmp.PushBack(value.Left)
			}
			if value.Right != nil {
				tmp.PushBack(value.Right)
			}
		}
		result = append(result, t)
		if tmp.Len() <= 0 {
			break
		}
		dataList.Init()
		for e := tmp.Front(); e != nil; e = e.Next() {
			dataList.PushBack(e.Value.(*TreeNode))
		}
		tmp.Init()
	}
	length := len(result)
	for i := 0; i < length / 2; i ++ {
		result[i], result[length - i - 1] = result[length - i - 1] , result[i]
	}
	return result
}

func main() {
	root := &TreeNode{Val:3}
	root.Left = &TreeNode{Val:9}
	root.Right = &TreeNode{Val:20}
	root.Right.Left = &TreeNode{Val:15}
	root.Right.Right = &TreeNode{Val:7}
	fmt.Println(levelOrderBottom(root))
}
