package main

import "fmt"

func lengthOfLastWord(s string) int {
	var worldLength int
	if len(s) <= 0 {
		return 0
	}
	length := len(s)
	isLast := false
	for i := length - 1 ; i >= 0 ;i -- {
		if s[i] != ' ' {
			isLast = true
			worldLength = worldLength + 1
		}else {
			if isLast {
				break
			}
		}
	}
	return worldLength
}

func main() {
	//s := "Hello World"
	s := "a"
	fmt.Println(lengthOfLastWord(s))
}
