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

func deleteDuplicates(head *ListNode) *ListNode {
    for e := head ; e != nil ;  {
        if e != nil && e.Next != nil && e.Val == e.Next.Val {
            e.Next = e.Next.Next
        }else {
            e = e.Next
        }
    }
    return head
}
func swapPairs(head *ListNode) *ListNode {
 if head == nil || head.Next == nil {
        return head
    }
    next := head.Next
    head.Next = swapPairs(next.Next)
    next.Next = head
    return next
}
func main() {
	nums := []int{1,1,2,3,3,4,5,6,6,6,6,6,7,8,9,10}
	node := create(nums)
        show(deleteDuplicates(node))
}
