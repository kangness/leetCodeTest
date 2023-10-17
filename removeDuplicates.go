package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	length := len(nums)
	for i := 0; i < len(nums)-2; i++ {
		for {
			if i >= length-2 {
				break
			}
			if nums[i] == nums[i+2] {
				length--
				for j := i; j < len(nums)-1; j++ {
					nums[j] = nums[j+1]
				}
			} else {
				break
			}
		}
	}
	return length
}

func main() {
	//nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	//nums := []int{1,1,2}
	//nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 8}
	nums := []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	fmt.Println(nums)
	fmt.Println("orign len:", len(nums))
	fmt.Println(removeDuplicates(nums))
	fmt.Println(nums)
}
