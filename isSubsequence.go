package main

import "fmt"

func isSubsequence2(s ,t string) bool {
	n, m := len(s), len(t)
	i, j := 0, 0
	for i < n && j < m {
		if s[i] == t[j] {
			i++
		}
		j++
	}
	return i == n
}
func isSubsequence(s string, t string) bool {
	if len(s) <= 0 {
		return true
	}
	if len(t) <= 0 {
		return false
	}
	length := len(t)
	slength := len(s) -1
	for m := 0; m < length; m ++ {
		n := m
		si := 0
		for {
			if t[n] == s[si] {
				si ++
				n ++
			}else {
				n ++
			}
			if n >= length {
				break
			}
			if si == slength && t[n] == s[si] {
				return true
			}
		}
	}
	return false
}

func main() {
	s := "axc"
	t := "ahbgdc"
	fmt.Println(isSubsequence(s, t))
}
