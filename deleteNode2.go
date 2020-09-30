package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

var head *ListNode
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


func deleteNode(node *ListNode) {
	if head == nil {
		return
	}
	h := &ListNode{}
	h.Next = head
	for e := h; e != nil ; e = e.Next {
		if e.Next != nil && e.Next.Val == node.Val {
			e.Next = e.Next.Next
			break
		}
	}
	return
}
