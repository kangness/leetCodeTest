package main

import (
	"github.com/kangness/leetCodeTest/node_list"
)


func reverseKGroup2(head *node_list.ListNode, k int) *node_list.ListNode {
	length := 0
	h := head
	next := head
	dummy := &node_list.ListNode{Next: head}
	for {
		if h == nil {
			break
		}
		if h != nil {
			length ++
			h = h.Next
		}
	}
	curr := head
	prev := dummy
	for  i := 0; i < length / k; i++ {
		for j := 0; j < k - 1; j++ {
			next = curr.Next
			curr.Next = next.Next
			next.Next = prev.Next
			prev.Next = next
		}
		prev = curr
		curr = prev.Next
	}
	return dummy.Next
}

func main() {
	//head = [1,2,3,4,5], k = 2
	tmp := []int{1,2,3,4,5}
	head := node_list.Create(tmp)
	node_list.ShowNodeList(head)
	node_list.ShowNodeList(reverseKGroup2(head,2))
}
