package main

import "fmt"

func missingNumber(nums []int) int {
	tmp := make([]int,len(nums) + 1)
	for _, v := range nums {
		tmp[v] = 1
	}
	for i := 0 ;i <= len(nums); i ++ {
		if tmp[i] != 1 {
			return i
		}
	}
	return 0
}

func main() {
	nums := []int{9,6,4,2,3,5,7,0,1}
	fmt.Println(nums)
	fmt.Println(missingNumber(nums))
}
