package main

import (
	"container/list"
	"fmt"
)
type Parenth struct {
	offset int
	value byte
}
func longestValidParentheses(s string) int {
	if  len(s) <= 1 {
		return 0
	}
	q := list.New()
	q.PushFront(&Parenth{offset: 0, value:s[0]})
	length := len(s)
	result := 0
	tmp := make([]bool, length)
	for i:= 1; i < length; i ++ {
		if q.Len() > 0 {
			last := q.Front().Value.(*Parenth)
			if last.value == '(' && s[i] == ')' {
				q.Remove(q.Front())
				tmp[last.offset] = true
				tmp[i] = true
				continue
			}
		}
		q.PushFront(&Parenth{offset: i,value:s[i]})
	}
	value := 0
	fmt.Println(tmp)
	for _, v := range tmp {
		if v {
			value ++
			if result < value {
				result = value
			}
		}else {
				value = 0
		}

	}
	return result
}

func main() {
//	s := ")()())"
//	s := "()(((((())))))"
s := "()(()"
	fmt.Println(longestValidParentheses(s))
}
