package main

import (
	"fmt"
	"math"
	"strconv"
)

func isSushu(n int) bool {
	dt := int(math.Sqrt(float64(n))) + 2
	for j := 2; j <= dt; j++ {
		if n%j == 0 {
			return false
		}
	}
	return true
}
func isHuiWen(i int) bool {
	tmp := strconv.Itoa(i)
	if len(tmp) <= 1 {
		return true
	}
	length := len(tmp)
	for i := 0; i < length/2; i++ {
		if tmp[i] != tmp[length-i-1] {
			return false
		}
	}
	return true
}
func primePalindrome(n int) int {
	if n == 2 || n == 3 || n == 5 || n == 7 || n == 11 {
		return n
	}
	if n == 1 {
		return 2
	}
	for i := n; i < math.MaxInt32; i++ {
		if !isHuiWen(i) {
			continue
		}
		if !isSushu(i) {
			continue
		}
		return i
	}
	return 0
}
func main() {
	n := 9989900
	fmt.Println(primePalindrome(n))
}
