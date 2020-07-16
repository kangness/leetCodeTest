package main

import (
    "fmt"
)

func maxSubArray(nums []int) int {
	if nums == nil || len(nums) <= 0 {
            return 0
        }
        length := len(nums)
        if length <= 1 {
            return nums[0]
        }
	dp := make([]int, length+1)
        dp[0] = nums[0]
        for i := 1; i < length; i ++ {
		if dp[i - 1] > 0 {
			dp[i] = dp[i -1 ]
		}
	}
	fmt.Println(dp)
	return 0
}
func main() {
    nums := []int{-2,1,-3,4,-1,2,1,-5,4}
    fmt.Println(maxSubArray(nums))
}
