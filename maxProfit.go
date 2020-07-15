package main

import (
	"fmt"
)

func maxProfit(nums []int) int {
	if nums == nil || len(nums) <= 1 {
		return 0
	}
	length := len(nums)
	result := 0
	for i := 0; i < length; i++ {
		for j := i; j < length; j++ {
			if nums[i] < nums[j] {
				if nums[j]-nums[i] > result {
					result = nums[j] - nums[i]
				}
			}
		}
	}
	return result
}
func main() {
	nums := []int{7, 1, 5, 4, 6, 4}
	fmt.Println(maxProfit(nums))
}
