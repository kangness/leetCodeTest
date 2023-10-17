package main

import "fmt"

func canJump(nums []int) bool {
	length := len(nums)
	dp := make([]int, length)
	// dp[i] 表示从第i位能否到达终点
	dp[length-1] = 1
	for i := length - 1; i >= 0; i-- {
		for j := i; j <= i+nums[i]; j++ {
			fmt.Println("before i: ", i, "j:  ", j, "num[i]:", nums[i], dp)
			if j > length {
				break
			}
			if j < length && dp[j] == 1 {
				dp[i] = 1
				break
			}
			fmt.Println("after i: ", i, "j:  ", j, "nums[i]:", nums[i], dp)
		}
	}
	if dp[0] == 1 {
		return true
	}
	return false
}

func main() {
	nums := []int{2, 3, 1, 1, 4}
	fmt.Println(canJump(nums))
}
