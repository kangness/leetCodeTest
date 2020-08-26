package main

import "fmt"

func maxSubArray2(nums []int)int {
	var dp []int
	if nums == nil || len(nums) <= 0 {
		return 0
	}
	length := len(nums)
	dp = make([]int, length)
	dp[0] = nums[0]
	result := nums[0]
	for i := 1; i < length; i ++ {
		if dp[i - 1] < 0 {
			dp[i] = nums[i]
		}else if dp[i -1] >= 0 {
			dp[i] = dp[i -1] + nums[i]
		}
		if result < dp[i] {
			result = dp[i]
		}
	}
	return result
}

func main() {
	nums := []int{-2,1,-3,4,-1,2,1,-5,4}
	fmt.Println(nums)
	fmt.Println(maxSubArray2(nums))
}
