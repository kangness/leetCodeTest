package main

import "fmt"

func removeElement(nums []int, val int) int {
	if nums == nil || len(nums) <= 0 {
		return 0
	}
	length := len(nums)
	var j int
	j = 0
	for i := 0; i < length; i ++ {
		if nums[i] != val {
			nums[j] = nums[i]
			j = j + 1
		}
	}
	return j
}

func main () {
	nums := [] int{1,2,3,4,5,6,7,8,10,1,2,3,4}
	fmt.Println(removeElement(nums, 1))
}