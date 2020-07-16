package main

import (
	"fmt"
)

/*
给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。

示例 1:

输入: coins = [1, 2, 5], amount = 11
输出: 3
解释: 11 = 5 + 5 + 1
示例 2:

输入: coins = [2], amount = 3
输出: -1
说明:
你可以认为每种硬币的数量是无限的。

解题思路：

使用动态规划的思想来解决问题.
11 = min{1+f(11-1),1+f(11-2),1+f(6)}
大问题的最优解可以根据小问题的最优解解答。
*/
func iconChange(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i <= amount; i++ {
		dp[i] = amount + 1
	}
	dp[0] = 0
	for i := 1; i <= amount; i++ {
		for _, c := range coins {
			if i >= c {
				if dp[i] > dp[i-c] {
					dp[i] = dp[i-c] + 1
				}
			}
		}
	}
	fmt.Println(dp)
	fmt.Println(dp[amount])
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]

}
func main() {
	coins := []int{1, 2, 5}
	amount := 11
	fmt.Println(iconChange(coins, amount))
}
