package node_list

func isPalindrome(head *ListNode) bool {
	n := 1
	sum := 0
	for e := head; e != nil ; e = e.Next {
		sum = n * e.Val + sum
		n = n + 1
	}
	for e := head; e != nil; e = e.Next {
		n = n + 1
		sum = sum - ( n * e.Val)
	}
	if sum == 0 {
		return true
	}
	return false
}


