package main

import "fmt"

func countSegments(s string) int {
	s = " " + s
	length := len(s)
	var result int
	for i := 1; i < length; i ++ {
		if s[i - 1] == ' ' && s[i] != ' ' {
			result++
		}
	}
	return result
}

func main() {
//	s := "Hello,     my name is John"
	s := "a"
	fmt.Println(countSegments(s))
}
