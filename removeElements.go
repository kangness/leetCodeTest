package main

import (
	"fmt"
//	"math/rand"
//	"time"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func Show(node *ListNode) {
	if node == nil {
		return
	}
	i := 0
	for e := node; e != nil; e = e.Next {
		fmt.Println(i, e.Val)
		i ++
	}
}

func Create(num []int)  *ListNode {
	if len(num) <= 0 {
		return nil
	}
	start := &ListNode{}
	node := start
	length := len(num)
	for i, n := range num {
		node.Next = &ListNode{}
		node.Val = n
		if i < length -1 {
			node = node.Next
		}
	}
	node.Next = nil
	return start
}



func removeElements(head *ListNode, val int) *ListNode {
	h := &ListNode{}
	h.Next = head
	cur := h
	for {
		if cur.Next == nil {
			break
		}
		if cur.Next != nil && cur.Next.Val == val {
			cur.Next = cur.Next.Next
		}else {
			cur = cur.Next
		}
	}
	return h.Next
}

func main() {
	var nums []int
	/*
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 20; i ++ {
		nums = append(nums, rand.Int() % 20)
	}
	 */
	nums = []int{1,2,3,3,3,3,3,4}
	fmt.Println(nums,3)
	head := Create(nums)
	//Show(removeElements(head, nums[3]))
	Show(removeElements(head, 3))
}
