package main
import (
"fmt"
)
type ListNode struct {
    Val int
    Next *ListNode
}

func show(node *ListNode) {
	i := 0
	for e := node; e != nil; e = e.Next {
		fmt.Println(i, e.Val)
		i ++
	}
}

func create(num []int)  *ListNode {
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
func deleteNode(head *ListNode, val int) *ListNode {
	if head == nil {
		return head
	}
	if head.Val == val {
		return head.Next
	}
	head.Next = deleteNode(head.Next, val)
	return head
}
func main() {
	nums := []int{1,2,3,4,5,6,7,8,9,10}
	node := create(nums)
	show(node)
}
