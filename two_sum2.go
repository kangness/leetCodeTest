package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)
	for index, n := range nums {
		numMap[n] = index
	}
	length := len(nums)
	for i := 0; i <= length-1; i++ {
		head := target - nums[i]
		index, ok := numMap[head]
		if ok && index != i {
			return []int{i, index}
		}

	}
	return []int{}
}

func main() {
	nums := []int{3, 2, 4}
	target := 6
	fmt.Println(twoSum(nums, target))
}
