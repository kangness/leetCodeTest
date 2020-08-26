package main
import (
	"fmt"
)
type ListNode struct {
	Val int
	Next *ListNode
}

func Show(node *ListNode) {
	return
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


func MergeTwoLists(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = MergeTwoLists(l1.Next,l2)
		return l1
	}else {
		l2.Next = MergeTwoLists(l1, l2.Next)
		return l2
	}
}

func MergeTwoList2(l1, l2 *ListNode) *ListNode {
	result := &ListNode{}
	start := result
	for {
		var tmp *ListNode
		if l1 == nil && l2 == nil {
			break
		}
		if l1 != nil && l2 != nil && l1.Val >= l2.Val {
			tmp = l2
			l2 = l2.Next
		}else if l1 != nil && l2 != nil && l1.Val < l2.Val {
			tmp = l1
			l1 = l1.Next
		}else if l1 != nil && l2 == nil {
			tmp = l1
			l1 = l1.Next
		}else if l1 == nil && l2 != nil {
			tmp = l2
			l2 = l2.Next
		}else {
			fmt.Println("9999999999999")
		}
		fmt.Println(l1, l2, result)
		if result != nil {
			result.Val = tmp.Val
			result.Next = &ListNode{}
			result = result.Next
		}
	}
	return start
}

func main () {
	v1 := []int{1,3,5,6,9}
	v2 := []int{1,2,3,4,5,6,7,8,9}
	l1 := Create(v1)
	l2 := Create(v2)
	Show(MergeTwoList2(l1, l2))
	fmt.Println("-------------")
	Show(MergeTwoLists(l1,l2))
}
