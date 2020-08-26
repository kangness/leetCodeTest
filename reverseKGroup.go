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

func reversePrint(head *ListNode) []int {
	var tmp []int
	for e := head; e != nil ; e = e.Next {
		tmp = append(tmp, e.Val)
	}
	length := len(tmp)
	var result []int
	for i := length - 1 ; i >= 0 ; i -- {
		result = append(result, tmp[i])
	}
	return result
}

func reverseKGroup(head *ListNode, n int) *ListNode {
    if head == nil || head.Next == nil {
        return head
    }
    next := head.Next
    i := 0
    var p *ListNode
    for e := head; e != nil ; e = e.Next {
	p = e.Next
        p.Next = 
    }
    head.Next = swapPairs(next.Next)
    next.Next = head
    return next
}
func main() {
	nums := []int{1,2,3,4,5,6,7,8,9,10}
	node := create(nums)
	show(node)
}
