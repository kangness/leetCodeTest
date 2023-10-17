package main

import (
	"fmt"
	"math"
)

func maxProfit(nums []int) int {
	if nums == nil || len(nums) <= 1 {
		return 0
	}
	length := len(nums)
	dp := make([]int, length)
	dp[0] = 0
	result := 0
	for i := 1; i < length; i++ {
		if nums[i] > nums[i-1] {
			dp[i] = dp[i-1] + nums[i] - nums[i-1]
		} else {
			dp[i] = dp[i-1]
		}
		if result < dp[i] {
			result = dp[i]
		}
	}
	return result
}

func maxProfit2(nums []int) int {
	minPrice := nums[0]
	ans := 0
	for i := 0; i < len(nums); i++ {
		if ans < nums[i]-minPrice {
			ans = nums[i] - minPrice
		}
		if nums[i] < minPrice {
			minPrice = nums[i]
		}
	}
	return ans
}

/*
给你一个整数数组 prices ，其中 prices[i] 表示某支股票第 i 天的价格。
在每一天，你可以决定是否购买和/或出售股票。你在任何时候 最多 只能持有 一股 股票。你也可以先购买，然后在 同一天 出售。
返回 你能获得的 最大 利润 。
*/

func maxProfit3(prices []int) int {
	length := len(prices)
	dp := make([]int, length)
	for i := 1; i < length; i++ {
		if prices[i] > prices[i-1] {
			dp[i] = prices[i] - prices[i-1]
		}
	}
	total := 0
	for i := 0; i < length; i++ {
		total = total + dp[i]
	}
	return total
}

/*
给定一个数组，它的第 i 个元素是一支给定的股票在第 i 天的价格。
设计一个算法来计算你所能获取的最大利润。你最多可以完成 两笔 交易。
注意：你不能同时参与多笔交易（你必须在再次购买前出售掉之前的股票）。
*/

func maxProfit4(prices []int) int {
	length := len(prices)
	dpMap := make([]int, length)
	for i := 0; i < length; i++ {
		minPrice := prices[0]
		ans := 0
		for j := 0; j <= i; j++ {
			if ans < prices[j]-minPrice {
				ans = prices[j] - minPrice
			}
			if minPrice > prices[j] {
				minPrice = prices[j]
			}
		}
		dpMap[i] = dpMap[i] + ans
		minPrice = prices[i]
		ans = 0
		for j := i; j < length; j++ {
			if prices[j]-minPrice > ans {
				ans = prices[j] - minPrice
			}
			if prices[j] < minPrice {
				minPrice = prices[j]
			}
		}
		dpMap[i] = dpMap[i] + ans
	}
	maxPrice := dpMap[0]
	for i := 0; i < length; i++ {
		if maxPrice < dpMap[i] {
			maxPrice = dpMap[i]
		}
	}
	return maxPrice
}
func maxProfit5(prices []int) int {
	buy1, buy2 := -math.MaxInt, -math.MaxInt
	sell1, sell2 := 0, 0
	for _, price := range prices {
		buy1 = int(math.Max(float64(buy1), float64(-price)))
		sell1 = int(math.Max(float64(sell2), float64(buy1+price)))
		buy2 = int(math.Max(float64(buy2), float64(sell1-price)))
		sell2 = int(math.Max(float64(sell2), float64(buy2+price)))
	}
	return sell2
}
func maxProfit6(k int, prices []int) int {
	var dp [][][]int
	length := len(prices)
	dp = make([][][]int, k+1)
	for i := 0; i < k; i++ {
		dp[i] = make([][]int, length)
		for j := 0; j < length; j++ {
			dp[i][j] = make([]int, 4)
			dp[1][i][0] = -prices[i]
		}
	}
	// dp[k][i][0] => 表示购买过k次股票的收益
	// dp[k][i][1] = > 表示出售过k次股票的收益
	for i := 1; i < length; i++ {
		for j := 1; j <= k; j++ {
			dp[j][i][0] = int(math.Max(float64(dp[j][i-1][0]), float64(dp[j][i][1]+prices[i])))
		}
	}

}
func main() {
	//nums := []int{7, 1, 5, 4, 6, 4}
	//nums := []int{7, 1, 5, 3, 6, 4}
	//nums := []int{1, 2, 3, 4, 5}
	//nums := []int{3, 3, 5, 0, 0, 3, 1, 4}
	//nums := []int{1, 2, 3, 4, 5}
	nums := []int{3, 3, 5, 0, 0, 3, 1, 4}
	//nums := []int{3, 3, 5, 0}
	fmt.Println(maxProfit5(nums))
}
