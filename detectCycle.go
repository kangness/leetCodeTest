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


func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	tmp := make(map[*ListNode]bool)
	p := head
	for {
		_, ok := tmp[p]
		if ok {
			return p
		}
		tmp[p] = true
		p = p.Next
		if p == nil {
			break
		}
	}
	return nil
}

func main () {
	h := []int{3,2,0,-4}
	head := Create(h)
	//head.Next.Next.Next = head.Next
	fmt.Println(isCycle(head))
}
