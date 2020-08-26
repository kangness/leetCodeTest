package node_list

func MergeTwoLists(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val > l2.Val {
		l1.Next = MergeTwoLists(l1.Next,l2)
		return l1
	}
	l2.Next = MergeTwoLists(l1,l2.Next)
	return l2
}

func MergeTwoList2(l1, l2 *ListNode) *ListNode {
	var result *ListNode
	for {
		var tmp *ListNode
		if l1 == nil && l2 == nil {
			break
		}
		if l1 != nil && l2 != nil && l1.Val >= l2.Val {
			tmp = l2

		}else if l1 != nil && l2 != nil && l1.Val < l2.Val {
			tmp = l1

		}else if l1 != nil && l2 == nil {
			tmp = l1
			l1 = l1.Next
		}else if l1 == nil && l2 != nil {
			tmp = l2
			l2 = l2.Next
		}
		if result != nil {
			result.Next = tmp
			tmp = tmp.Next
		}else {
			result = tmp
		}
	}
	return result
}
