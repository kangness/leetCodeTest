package main

import "fmt"

func matrixReshape(nums [][]int, r, c int) [][]int {
	var lr, lc , total int
	lr = len(nums)
	if lr > 0 {
		lc = len(nums[0])
	}
	if r * c != lr * lc {
		return nums
	}
	result := make([][]int, r)
	for i := 0; i < r;i++ {
		result[i] = make([]int, c)
	}
	total = r * c
	for i := 0; i < total; i ++ {
		result[i/c][i%c] = nums[i/lc][i%lc]
	}
	return result
}

func main() {
	nums := [][]int{{1,2},{3,4}}
	r := 2
	c := 4
	result := matrixReshape(nums, r,c)
	for _, m := range result {
		fmt.Println(m)
	}
}
