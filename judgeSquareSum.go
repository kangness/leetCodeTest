package main

import (
	"fmt"
	"math"
)

func judgeSquareSum(c int) bool {
	if c == 0 {
		return true
	}
	for i := 0; i <= c; i ++ {
		if i * i > c {
			break
		}
		for j := 0 ;j <= c; j ++ {
			tmp := i * i + j* j
			if tmp == c {
				fmt.Println(i, j)
				return true
			}
			if j * j > c {
				break
			}

		}
	}
	return false
}

func judgeSquareSum2(c int) bool {
	if c <= 2{
		return true
	}
	m := int(math.Sqrt(float64(c)))
	var max, min int
	max = m
	for {
		if max < min {
			break
		}
		tmp := max * max + min * min
		if tmp == c {
			return true
		}
		if tmp > c {
			max = max - 1
		}else {
			min = min + 1
		}
	}
	return false
}

func main() {
	c := 1000
	fmt.Println(judgeSquareSum(c))
	fmt.Println(judgeSquareSum2(c))
}
