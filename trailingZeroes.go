package main

import "fmt"

func trailingZeroes(n int) int {
	five := 0
	buffer := make(map[int]int)
	buffer[5] = 1
	for i := 1; i <= n ; i ++ {
		if i >= 5 {
			tmp := i
			for {
				if tmp % 5 == 0 {
					value, _ := buffer[tmp/5]
					five = five + value + 1
					buffer[tmp] = value + 1
				}
				break
			}
		}
	}
	return five
}
func trailingZeroes2(n int) int {
	if n < 5 {
		return 0
	}
	five := 0
	for {
		five = five + n / 5
		n = n / 5
		if n <4 {
			break
		}
	}
	return five
}
func main() {
	n := 18085483
	fmt.Println(trailingZeroes(n))
}
