package node_list

import "fmt"

// 创建联表
func CreateListNodeByArray(vals []int) *ListNode {
	head := &ListNode{}
	tmp := head
	for _, v := range vals {
		tmp.Val = v
		tmp.Next = &ListNode{}
		tmp = tmp.Next
		tmp.Next = nil
	}
	return head
}
// 打印联表
func ShowNodeList(head *ListNode) []int{
	h := head
	var result []int
	for {
		if h == nil {
			break
		}
		if h != nil {
			result = append(result,h.Val)
		}
		h = h.Next
	}
	fmt.Println(result)
	return result
}
