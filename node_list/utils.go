package node_list
import (
	"fmt"
)
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

