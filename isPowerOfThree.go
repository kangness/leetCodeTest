package main

import "fmt"

func isPowerOfThree(n int) bool {
	if n == 0 {
		return false
	}
	for {
		if n == 1 {
			return true
		}
		if n <= 0 {
			return false
		}
		if n % 3 != 0 {
			return false
		}
		n = n  / 3
	}
	return false
}

func main() {
	n := 27
	fmt.Println(isPowerOfThree(n))
}
