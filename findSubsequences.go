package main

import "fmt"

func findSubsequences(nums []int) [][]int {
	length := len(nums)
	for i := 0; i < length; i ++ {
		for j := i; j < length; j ++ {
			if nums[i] > nums[j] {
				nums[i],nums[j] = nums[j], nums[i]
			}
		}
	}

	fmt.Println(nums)
	return  nil
}

func main() {
	nums := []int{4, 6, 7, 7}
	fmt.Println(findSubsequences(nums))
}
