package main
import (
		"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    result := l1
    if l1 == nil {
        return l2
    }
    if l2 == nil {
        return l1
    }
    start := l1
    for {
        if l1 == nil && l2 != nil {
            start.Next = l2
            l2 = l2.Next
        }else if l2 == nil && l1 != nil {
            start.Next = l1
            l1 = l1.Next
        }else if l2 == nil && l1 == nil {
            break
        }else if l1.Val <= l2.Val {
            start.Next = l1
            l1 = l1.Next
        }else {
            start.Next = l2
            l2 = l2.Next
        }
        start = start.Next
    }
    return result
}

func main() {
   num1 := []int{2,4,3} 
   num2 := []int{5,6,4}
   addTwoNumbers(num1,num2)   
}
