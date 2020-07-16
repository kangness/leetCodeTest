package main

import (
	"fmt"
)

func maxSubArray(nums []int) int {
	if nums == nil || len(nums) <= 0 {
		return 0
	}
	length := len(nums)
	dp := make([]int, length+1)
	dp[0] = nums[0]
	res := nums[0]
	for i := 1; i < length; i++ {
		if dp[i-1]+nums[i] > nums[i] {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
		if res < dp[i] {
			res = dp[i]
		}
	}
	fmt.Println(dp)
	return res
}
func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(maxSubArray(nums))
}
