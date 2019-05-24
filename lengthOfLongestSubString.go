package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	ss := []byte(s)
	length := len(ss)
	if length <= 0 {
		return 0
	}
	if length <= 1 {
		return 1
	}
        temp := make(map[byte]int)
	max := 0
	start := -1
	for i , s := range ss{
		offset, ok := temp[s]	
		if ok {
			if i - offset <= 1 {
				start = i
			}
			if start < offset {
				start = offset
			}
			if i - start >= max {
				max = i - start
			}
		}
		temp[s] = i
	}
	if length  - 1 - start > max {
		max = length - start
	}
	return max
}
func main() {
	s := "abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
}
