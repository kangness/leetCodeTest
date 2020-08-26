package main

import "fmt"

func rotate(nums []int, k int) {
	length := len(nums)
	result := make([]int, length)
	for i := 0; i < length; i ++ {
		offset := (i + k) % length
		result[offset] = nums[i]
	}
	for i := 0; i < length; i ++ {
		nums[i] = result[i]
	}
}

func main() {
	nums := []int{1,2,3,4,5,6,7}
	k := 3
	rotate(nums, k)
	fmt.Println(nums)
}
