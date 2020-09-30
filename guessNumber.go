package main

import "fmt"

var guessNum int = 6
func guess(n int) int {
	if n > guessNum {
		return -1
	}else if n < guessNum {
		return 1
	}
	return 0
}
func guessNumber(n int) int {
	max := n
	min := 0
	for {
		mid := (max - min) / 2 + min
		res := guess(mid)
		if res == 0 {
			return mid
		}
		if res < 0 {
			max = mid - 1
		}else {
			min = mid + 1
		}
		if min > max {
			return 0
		}
	}
	return 0
}

func main () {
	n := 1000
	fmt.Println(guessNumber(n))
}
