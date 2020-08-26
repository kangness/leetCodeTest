package main

import (
	"fmt"
	"strings"
)

func isPalindromeStr(s string) bool {
	if len(s) <= 1 {
		return true
	}
	length := len(s)
	s = strings.ToLower(s)
	ss := []byte(s)
	n := 0
	m := length - 1
	for {
		for {
			if n > m || n >= length {
				return true
			}
			if (ss[n] >= 'a' && ss[n] <= 'z' ) || (ss[n] >= '0' && s[n] <= '9') {
				break
			}
			n ++
		}
		for {
			if m < n  || m < 0{
				return true
			}
			if (ss[m] >= 'a' && ss[m] <= 'z' ) || (ss[m] >= '0' && s[m] <= '9') {
				break
			}
			m --
		}
		if ss[m] != ss[n] {
			return false
		}
		n ++
		m --
		if m <= n {
			break
		}
	}
	return true
}

func main() {
//	s := "A man, a plan, a canal: Panama"
	s := "a."
	fmt.Println(isPalindromeStr(s))
}
