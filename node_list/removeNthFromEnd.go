package node_list


func removeNthFromEnd(head *ListNode, n int) *ListNode {
	i := 0
	var  cur *ListNode
	for e := head; e != nil ; e = e.Next {
		i ++
		if cur != nil {
			cur = cur.Next
		}
		if i == n - 1 {
			cur = e
		}
	}
	if cur != nil && cur.Next != nil{
		cur.Next = cur.Next
	}
	return nil
}
