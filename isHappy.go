package main

import "fmt"

func isHappy(n int) bool {
	for {
		tmp := 0
		for {
			nn := n % 10
			tmp = tmp + nn * nn
			n = n / 10
			if n <= 0 {
				break
			}
		}
		fmt.Println(tmp)
		if tmp == 1 {
			return true
		}
		if tmp <= 9 {
			return false
		}
		n = tmp
	}
	return true
}

func main() {
	n := 1111111
	fmt.Println(isHappy(n))
}
