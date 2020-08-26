package main

import (
	"fmt"
)

func maxProfit(nums []int) int {
	if nums == nil || len(nums) <= 1 {
		return 0
	}
	length := len(nums)
	dp := make([]int, length)
	dp[0] = 0
	result := 0
	for i := 1; i < length; i ++ {
		if nums[i] > nums[i -1] {
			dp[i] = dp[i - 1] + nums[i] - nums[i - 1]
		}else {
			dp[i] = dp[i - 1]
		}
		if result < dp[i] {
			result = dp[i]
		}
	}
	return result
}

func main() {
	//nums := []int{7, 1, 5, 4, 6, 4}
	nums := []int {7,1,5,3,6,4}
	fmt.Println(maxProfit(nums))
}
