package main

import "fmt"

func moveZeroes(nums []int) {
	length := len(nums)
	for i := 0; i < length; i ++ {
		if nums[i] != 0 {
			continue
		}
		for j := i; j < length; j ++ {
			if nums[j] != 0 {
				nums[i], nums[j] = nums[j], nums[i]
				break
			}
		}
	}
}

func main() {
	nums := []int{0,1,0,3,12}
	fmt.Println(nums)
	moveZeroes(nums)
	fmt.Println(nums)
}
