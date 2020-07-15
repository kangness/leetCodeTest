package main

/*
给定一个整数数组 nums ，找到一个具有最大和的连续子数组（子数组最少包含一个元素），返回其最大和。

示例:

输入: [-2,1,-3,4,-1,2,1,-5,4],
输出: 6
解释: 连续子数组 [4,-1,2,1] 的和最大，为 6。
进阶:

如果你已经实现复杂度为 O(n) 的解法，尝试使用更为精妙的分治法求解。

解题思路：
下面介绍动态规划的做法，复杂度为 O(n)。
步骤 1：令状态 dp[i] 表示以 A[i] 作为末尾的连续序列的最大和（这里是说 A[i]必须作为连续序列的末尾）。
步骤 2：做如下考虑：因为 dp[i] 要求是必须以 A[i] 结尾的连续序列，那么只有两种情况：
这个最大和的连续序列只有一个元素，即以 A[i] 开始，以 A[i] 结尾。
这个最大和的连续序列有多个元素，即从前面某处 A[p] 开始 (p<i)，一直到 A[i] 结尾。

对第一种情况，最大和就是 A[i] 本身。
对第二种情况，最大和是 dp[i-1]+A[i]。

于是得到状态转移方程：
dp[i] = max{A[i], dp[i-1]+A[i]}

*/
import (
	"fmt"
)

func sumMax(nums []int) int {
	if len(nums) <= 0 {
		return 0
	}
	lengthNum := len(nums)
	dp := make([]int, lengthNum)
	dp[0] = nums[0]
	result := 0
	for i := 1; i < lengthNum; i++ {
		if dp[i-1]+nums[i] < nums[i] {
			dp[i] = nums[i]
		} else {
			dp[i] = dp[i-1] + nums[i]
		}
		if result < dp[i] {
			result = dp[i]
		}
	}
	fmt.Println(dp)
	return result
}
func main() {
	nums := []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}
	fmt.Println(sumMax(nums))
}
