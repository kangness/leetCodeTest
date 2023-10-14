package main

import (
	"fmt"
	"math"
)

func divide(dividend int, divisor int) int {
	if dividend == 0 {
		return 0
	}
	if divisor == 1 {
		return dividend
	}
	if divisor == -1 {
		if dividend > math.MinInt32 {
			return -dividend
		} // 只要不是最小的那个整数，都是直接返回相反数
		return math.MaxInt32 // 是最小的那个，那就返回最大的整数
	}

	sign := 1
	if (dividend < 0 && divisor > 0) || (dividend > 0 && divisor < 0) {
		sign = -1
	}
	if dividend < 0 {
		dividend = -dividend
	}
	if divisor < 0 {
		divisor = -divisor
	}
	res := divideHelper(dividend, divisor)
	if sign > 0 {
		if res > math.MaxInt32 {
			return math.MaxInt32
		} else {
			return res
		}
	}
	return -res
}

func divideHelper(a int, b int) int {
	if a < b {
		return 0
	}
	count := 1
	tb := b
	for tb+tb <= a {
		count = count + count
		tb = tb + tb
	}
	return count + divideHelper(a-tb, b)
}

func main() {
	dividend := -2147483648
	divisor := 1
	fmt.Println(divide(dividend, divisor))
}
