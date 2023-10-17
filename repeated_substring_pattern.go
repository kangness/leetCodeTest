package main

import (
	"fmt"
	"strings"
)

func repeatedSubstringPattern(s string) bool {
	tmp := s + s
	return strings.Contains(tmp, tmp[1:len(tmp)-1])
}

func main() {
	s := "ababba"
	fmt.Println(repeatedSubstringPattern(s))
}
