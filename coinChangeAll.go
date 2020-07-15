package main

import (
	"fmt"
)

/*
给定不同面额的硬币和一个总金额。写出函数来计算可以凑成总金额的硬币组合数。假设每一种面额的硬币有无限个。

注意: 你可以假设

0 <= amount (总金额) <= 5000
1 <= coin (硬币面额) <= 5000
硬币种类不超过500种
结果符合32位符号整数
示例 1:

输入: amount = 5, coins = [1, 2, 5]
输出: 4
解释: 有四种方式可以凑成总金额:
5=5
5=2+2+1
5=2+1+1+1
5=1+1+1+1+1
*/
func iconChangeAll(coins []int, amount int) int {
	dp := make([]int, amount+1)
	for i := 0; i <= amount; i++ {
		dp[i] = 0
	}
	dp[0] = 1
	for _, c := range coins {
		for i := c; i <= amount; i++ {
			dp[i] = dp[i] + dp[i-c]
		}
	}
	fmt.Println(dp)
	fmt.Println(dp[amount])
	return dp[amount]
}

func main() {
	coins := []int{5, 2, 1}
	amount := 5
	fmt.Println(iconChangeAll(coins, amount))
}
