package main

import (
	"fmt"
)

/*
你是一个专业的小偷，计划偷窃沿街的房屋。每间房内都藏有一定的现金，影响你偷窃的唯一制约因素就是相邻的房屋装有相互连通的防盗系统，如果两间相邻的房屋在同一晚上被小偷闯入，系统会自动报警。

给定一个代表每个房屋存放金额的非负整数数组，计算你在不触动警报装置的情况下，能够偷窃到的最高金额。

解题思路：
下面介绍动态规划的做法，复杂度为 O(n)。
步骤 1：令状态 dp[i] 表示以 A[i]作为被劫持房间的最大值。
步骤 2：做如下考虑：因为 dp[i]代表劫持到第i个房间的最打值，那么只有两种情况：
劫持该房间。
不劫持到该房间

对第一种情况，最大值就是 dp[i-2]+A[i] 本身。
对第二种情况，最大和是 dp[i-1]。

于是得到状态转移方程：
dp[i] = max{dp[i-2]+A[i], dp[i-1]}
*/
func rob(nums []int) int {
	if nums == nil || len(nums) <= 0 {
		return 0
	}
	length := len(nums)
	if length == 1 {
		return nums[0]
	}
	if length == 2 {
		if nums[0] > nums[1] {
			return nums[0]
		} else {
			return nums[1]
		}
	}
	dp := make([]int, length)
	dp[0] = nums[0]
	if nums[0] > nums[1] {
		dp[1] = nums[0]
	} else {
		dp[1] = nums[1]
	}
	for i := 2; i < length; i++ {
		if dp[i-2]+nums[i] > dp[i-1] {
			dp[i] = dp[i-2] + nums[i]
		} else {
			dp[i] = dp[i-1]
		}
	}
	fmt.Println(dp)
	return dp[length-1]
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(rob(nums))
}
