package main

import (
	"fmt"
	"container/list"
)

func isValid(s string) bool {
	if len(s) <= 1 {
		return false
	}
	if len(s) % 2 != 0 {
		return false
	}
	l := list.New()
	length := len(s)
	for i := 0 ; i < length; i ++ {
		if isSym(l,s[i]) {
			l.Remove(l.Back())
		}else {
			l.PushBack(s[i])
		}
	}
	if l.Len() <= 0 {
		return true
	}
	return false
}

func isSym(l *list.List, c2 byte) bool {
	if l == nil || l.Len() <= 0 || l.Front() == nil {
		return false
	}
	c1 := l.Back().Value.(byte)
	if (c1 == '(' && c2 == ')') || (c1 == '[' && c2 == ']') || (c1 == '{' && c2 == '}') {
		return true
	}
	return false
}

func main() {
	s := "([])"
	fmt.Println(isValid(s))
}
