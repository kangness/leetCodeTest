package main

import "fmt"

func maxPower(s string) int {
	length := len(s)
	if length <= 0 {
		return 0
	}
	if length <= 1 {
		return 1
	}
	result := 1
	tmp := 1
	for i := 1;i < length; i ++ {
		if s[i - 1] == s[i] {
			tmp = tmp + 1
		}else {
			tmp = 1
		}
		if result < tmp {
			result = tmp
		}
	}
	return result
}

func main() {
//	s := "leetcode"
//	s := "abbcccddddeeeeedcba"
//	s := "hooraaaaaaaaaaay"
	s := "tourist"
	fmt.Println(maxPower(s))
}
