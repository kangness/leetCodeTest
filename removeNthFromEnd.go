package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

func Show(node *ListNode) {
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



func removeNthFromEnd(head *ListNode, n int) *ListNode {
	i := 0
	var  cur *ListNode
	for e := head; e != nil ; e = e.Next {
		i ++
		if cur != nil {
			cur = cur.Next
		}
		if i == n + 1 {
			cur = head
		}
	}
	if cur != nil {
		if cur.Next != nil {
			cur.Next = cur.Next.Next
		}
	}
	if i == n {
		return head.Next
	}
	return head
}

func main() {
	nums := []int{1,2,3,4,5,6,7,8,9}
	head := Create(nums)
	Show(removeNthFromEnd(head, 9))
}
