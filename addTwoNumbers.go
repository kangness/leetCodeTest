package main

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func Show(node *ListNode) {
	for e := node; e != nil; e = e.Next {
		fmt.Print(e.Val)
	}
	fmt.Println()
}

func Show2(tag string, node *ListNode) {
	var tmp []int
	for e := node; e != nil; e = e.Next {
		tmp = append(tmp, e.Val)
	}
	fmt.Println(tag, tmp)
}

// func addTwoNumbers(l1 , l2 *ListNode) *ListNode {
func addTwoNumbers(num1, num2 []int) *ListNode {
	/*
	    var num1, num2,num3 []int
	    m := l1
	    n := l2
	    for {
	   	if m != nil {
	            num1 = append(num1, m.Val)
	            m = m.Next
	        }
	        if n != nil {
	            num2 = append(num2, n.Val)
	            n = n.Next
	        }
	        if m == nil && n == nil {
	            break
	        }
	    }
	*/
	var num3 []int
	length1 := len(num1)
	length2 := len(num2)
	if length1 < length2 {
		num1, num2 = num2, num1
		length1, length2 = length2, length1
	}
	num3 = append(num3, 0)
	num3 = append(num3, num1...)
	num1 = num3
	result := make([]int, 0, length1+1)
	number := 0
	for i := 0; i <= length2; i++ {
		if i < length2 {
			number = num1[length1-i] + num2[length2-i-1]
		} else {
			number = num1[length1-i]
		}
		if number > 9 {
			num1[length1-i-1] = num1[length1-i-1] + 1
		}
		result = append(result, number%10)
	}
	length := len(result)
	m := &ListNode{}
	res := m
	for i, n := range result {
		if i == length-1 && n == 0 {
			continue
		}
		m.Val = n
		m.Next = &ListNode{}
		m = m.Next
	}
	return res
}

func addTwoNumbers3(l1 *ListNode, l2 *ListNode) *ListNode {
	l1Num := 0
	l2Num := 0
	t1, t2 := l1, l2
	var l1a, l2a []int
	Show(l1)
	Show(l2)

	for {
		if t1 != nil {
			l1a = append(l1a, t1.Val)
			t1 = t1.Next
		} else {
			break
		}
	}
	for {
		if t2 != nil {
			l2a = append(l2a, t2.Val)
			t2 = t2.Next
		} else {
			break
		}
	}
	for i := len(l1a) - 1; i >= 0; i-- {
		l1Num = l1Num*10 + l1a[i]
	}
	for i := len(l2a) - 1; i >= 0; i-- {
		l2Num = l2Num*10 + l2a[i]
	}
	total := l2Num + l1Num
	fmt.Println(total)
	result := &ListNode{}
	tmp := result
	for {
		tmp.Next = &ListNode{}
		tmp.Val = total % 10
		total = total / 10
		if total <= 0 {
			tmp.Next = nil
			break
		}
		tmp = tmp.Next
	}
	return result
}
func Create(num []int) *ListNode {
	if len(num) <= 0 {
		return nil
	}
	start := &ListNode{}
	node := start
	length := len(num)
	for i, n := range num {
		node.Next = &ListNode{}
		node.Val = n
		if i < length-1 {
			node = node.Next
		}
	}
	node.Next = nil
	return start
}

func addTwoNumbers2(l1, l2 *ListNode) *ListNode {
	result := &ListNode{}
	tmp := 0
	for {
		t := 0
		if l1 == nil && l2 == nil {
			break
		}
		if l1 != nil {
			t = l1.Val + t
		}
		if l2 != nil {
			t = l2.Val + t
		}
		result.Val = t%10 + tmp
		tmp = t / 10
	}
	return nil
}

// 给定一个已排序的链表的头 head ， 删除原始链表中所有重复数字的节点，只留下不同的数字 。返回 已排序的链表 。
// 输入：head = [1,2,3,3,4,4,5]
// 输出：[1,2,5]
func deleteDuplicatesMine(head *ListNode) *ListNode {
	dup := &ListNode{Val: 0, Next: nil}
	cur := dup
	dup.Next = head
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			x := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == x {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}
	return dup.Next
}

func CreateRandList(n int) *ListNode {
	rand.Seed(time.Now().UnixNano())
	head := &ListNode{}
	head.Val = rand.Intn(1000)
	h := head
	for i := 0; i < n; i++ {
		h.Val = rand.Intn(1000)
		if i < n-1 {
			h.Next = &ListNode{}
			h = h.Next
		}
	}
	return head
}
func sortList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	var tmp []int
	for h := head; h != nil; h = h.Next {
		tmp = append(tmp, h.Val)
	}
	sort.Ints(tmp)
	result := &ListNode{}
	h := result
	for i := 0; i < len(tmp); i++ {
		h.Val = tmp[i]
		h.Next = nil
		if i < len(tmp)-1 {
			h.Next = &ListNode{}
			h = h.Next
		}
	}
	return result
}

func sortList2(head *ListNode) *ListNode {
	for tmp := head; tmp != nil; tmp = tmp.Next {
		for h := tmp.Next; h != nil; h = h.Next {
			if tmp.Val > h.Val {
				tmp.Val, h.Val = h.Val, tmp.Val
			}
		}
	}
	return head
}
func sortList3(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	if head.Next.Next == nil {
		if head.Val > head.Next.Val {
			var pre *ListNode
			cur := head
			for cur != nil {
				nextTmp := cur.Next
				cur.Next = pre
				pre = cur
				cur = nextTmp
			}
			return pre
		}
		return head

	}
	left := &ListNode{}
	ll := left
	right := &ListNode{}
	lr := right
	size := 0
	for h := head; h != nil; h = h.Next {
		size++
	}
	size = size / 2
	num := 0
	i := 0
	for h := head; h != nil; h = h.Next {
		i++
		if i >= size {
			num = h.Val
			break
		}
	}
	for h := head; h != nil; h = h.Next {
		if h.Val < num {
			ll.Next = &ListNode{}
			ll.Next.Val = h.Val
			ll = ll.Next
		} else {
			lr.Next = &ListNode{}
			lr.Next.Val = h.Val
			lr = lr.Next
		}
	}
	Show2("ll:", left.Next)
	Show2("lr:", right.Next)
	left = sortList3(left.Next)
	right = sortList3(right.Next)
	return mergeList(left, right)
}

func mergeTwoSortLists2(l1, l2 *ListNode) *ListNode {
	var head, tail *ListNode
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		head, tail, l1.Next = l1, l1, l1.Next
	} else {
		head, tail, l2 = l2, l2, l2.Next
	}
	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val {
			tail.Next, l1 = l1, l1.Next
		} else {
			tail.Next, l2 = l2, l2.Next
		}
		tail = tail.Next
	}
	if l1 != nil {
		tail.Next = l1
	}
	if l2 != nil {
		tail.Next = l2
	}
	return head
}
func mergeTwoSortLists(l1, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoSortLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoSortLists(l1, l2.Next)
		return l2
	}

	return nil
}
func mergeList(left, right *ListNode) *ListNode {
	if left == nil {
		return right
	}
	if right == nil {
		return left
	}
	h := left
	for ; h.Next != nil; h = h.Next {
	}
	h.Next = right
	return left

}
func ListFanzhuan(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		nextTmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = nextTmp
	}
	return pre
}
func main() {
	/*
		num1 := []int{2, 4, 3}
		num2 := []int{5, 6, 4}
		fmt.Println(num1, num2)
	*/
	//Show(addTwoNumbers3(Create(num1), Create(num2)))
	/*
		head := []int{1, 1, 1, 2, 3}
		Show(deleteDuplicatesMine(Create(head)))
	*/
	//l1 := sortList(CreateRandList(5))
	//l2 := sortList(CreateRandList(5))

	//sl2 := CreateRandList(3)

	//Show2("show sl:", sl)
	//Show2("show sortList:", sortList(sl))
	//Show2("merge two sort list", mergeTwoSortLists2(l1, l2))
	//Show2("sort list 2:", sortList3(CreateRandList(10)))
	l := CreateRandList(10)
	Show2("orgin:", l)
	Show2("swap:", sortList3(l))
}
