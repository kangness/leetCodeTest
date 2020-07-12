package main

import (
	"fmt"
	)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	var s []byte
	y := x
	for {
		 s = append(s, byte(y % 10) + '0')
		 y = y / 10
		 if y <= 0 {
		 	break
		 }
	}
	length := len(s)
	for i := 0; i < length / 2; i ++ {
		if s[i] != s[length - 1 - i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isPalindrome(1))
}
