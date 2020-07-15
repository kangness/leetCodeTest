package main

/*
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例:

输入: [-2,1,-3,4,-1,2,1,-5,4],
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
*/
import (
	"fmt"
)

func maxSubArray(a []int) int {
	length := len(a)
	b := make([]int, length, length)
	b[0] = a[0]
	result := a[0]
	for i := 1; i < length; i++ {
		if a[i] >= a[i]+b[i-1] {
			b[i] = a[i]
		} else {
			b[i] = a[i] + b[i-1]
		}
		if result < b[i] {
			result = b[i]
		}
	}
	return result
}
func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}
