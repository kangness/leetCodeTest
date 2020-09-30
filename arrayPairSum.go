package main

import (
	"fmt"
	"sort"
)

func arrayPairSum(nums []int) int {
	/*
	length := len(nums)
	for i := 0; i < length; i ++ {
		for j := i; j < length; j ++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	 */
	sort.Ints(nums)
	length := len(nums)
	var result int
	for i := 0; i < length; i = i + 2 {
		result = result + nums[i]
	}
	return result
}

func main() {
	nums := []int{1,4,3,2}
	fmt.Println(arrayPairSum(nums))
}
