package main
import (
		"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
    result := lists[0]
    for i := 1; i < len(lists); i++ {
        result = mergeTwoList(result, lists[i])
	show(result)
    }
    return result
}

func mergeTwoList(l1, l2 *ListNode) *ListNode {
    result := &ListNode{}
    for {
        if (l2 == nil && l1 != nil) || l1.Val <= l2.Val {
            result.Val = l1.Val
            l1 = l1.Next
        }
        if (l1 == nil && l2 != nil) || l1.Val > l2.Val {
            result.Val = l2.Val
            l2 = l2.Next
        }
        if l1 != nil && l2 != nil {
            result.Next = &ListNode{}
            result = result.Next
        }else {
             break
        }
    }
    return result
}
func show(lists *ListNode) {
    for {
	if lists != nil {
		fmt.Printf(lists.Val, " ")
		lists = lists.Next
	}else {
		break
	}
    }
}
func main() {
   var input []*ListNode 
    
   mergeKLists()   
}
