package main

import "fmt"

func firstUniqChar(s string) int {
	size := 26
	m := make([]int, size)
	n := make([]int, size)
	for i := 0; i < size; i ++ {
		n[i] = -1
	}
	for index, j := range s {
		offset := j - 'a'
		m[offset] ++
		n[offset] = index
	}
	min := len(s) + 1
	for i := 0; i < size; i ++ {
		if m[i] == 1 && n[i] != -1 && min > n[i] {
			min = n[i]
		}
	}
	if min >= len(s) + 1 {
		return -1
	}
	return min
}

func main() {
	s := "loveleetcode"
	fmt.Println(firstUniqChar(s))
}

