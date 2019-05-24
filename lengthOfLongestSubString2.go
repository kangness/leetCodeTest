package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	ss := []byte(s)
	if len(ss) <= 1 {
		return len(ss)
	}
	length := len(ss)
	temp := make(map[byte]int)
	start, max, old := 0, 0, 0
	for {
		if start >= length {
			break
		}
		offset, ok := temp[ss[start]]
		if !ok {
			temp[ss[start]] = start
		} else {
			delete(temp,ss[start])
			old = start
			start = offset
		}
		start++
		if start-old > max {
			max = start - old
		}
	}
	return max
}

func main() {
	s := "au"
	fmt.Println(s, lengthOfLongestSubstring(s))
}
