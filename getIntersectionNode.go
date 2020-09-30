package main

type ListNode struct {
	   Val int
	    Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if (headA == nil || headB == nil ) {
		return nil
	}
	p1 := headA
	p2 := headB
	for {
		if p1 == p2 {
			break
		}
		if p1 != nil {
			p1 = p1.Next
		}else {
			p1 = headB
		}
		if p2 != nil {
			p2 = p2.Next
		}else {
			p2 = headA
		}
	}
	return p1
}