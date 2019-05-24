package main
import (
		"fmt"
)

type ListNode struct {
	Val int
	Next *ListNode
}
//func addTwoNumbers(l1 , l2 *ListNode) *ListNode {
func addTwoNumbers(num1,num2 []int) *ListNode {
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
		number = num1[length1 - i] + num2[length2 - i - 1]
	}else {
		number = num1[length1 - i]
	}
        if number > 9 {
            num1[length1 - i - 1] = num1[length1 - i - 1] + 1
	}
        result = append(result, number % 10)
    }
    length := len(result)
    m := &ListNode{}
    res := m
    for i , n := range result {
        if i == length -1 && n == 0 {
		continue
	}
	m.Val = n
	m.Next = &ListNode{}
        m = m.Next
    }
	return res
}

func main() {
   num1 := []int{2,4,3} 
   num2 := []int{5,6,4}
   addTwoNumbers(num1,num2)   
}
