package main

import "fmt"

func titleToNumber(s string) int {
	ss := []byte(s)
	if  len(ss) <= 0 {
		return 0
	}
	length := len(ss)
	offset := 1
	result := 0
	for i := length - 1; i >= 0; i -- {
		result = result + offset * (int(ss[i] - 'A') + 1)
		offset = offset * 26
	}
	return result
}

func main() {
	s := "ZY"
	fmt.Println(titleToNumber(s))
}
