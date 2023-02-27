package main

import "fmt"

func minimumSwap(s1 string, s2 string) int {
	var x, y ,length int
	length = len(s1)
	for i := 0; i < length; i ++ {
		if s1[i] != s2[i] {
			if s1[i] == 'x' {
				x ++
			}else {
				y ++
			}
		}
	}
	if (x + y) % 2 == 1 {
		return -1
	}
	return (x+1) / 2 + (y + 1) /2
}
func main() {
	s1 := "yxxx"
	s2 := "xyyy"
	fmt.Println(minimumSwap(s1,s2))
}
