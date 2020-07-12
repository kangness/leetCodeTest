package main

import (
	"fmt"
)

func permute(nums []int) [][]int {
	var result [][]int
	length := len(nums)
	if length == 0 {
		return nil
	}
	if length == 1 {
		return [][]int{nums}
	}
	temp := permute(nums[:length - 1])
	for _, t := range temp {
		tl := len(t)
		for i := 0; i <= tl; i ++ {
			one := []int{}
			one = append(one,t[:i]...)
			one = append(one, nums[length - 1])
			one = append(one, t[i:tl]...)
			result = append(result, one)
		}
	}
	return result
}

func permutemn(nums []int,n int) [][]int {
	var result [][]int
	if n == 1 {
		for i := 0; i < len(nums); i ++ {
			result = append(result,[]int{nums[i]})
		}
		return result
	}
	temp := permutemn(nums,n - 1)
	for _, t := range temp {
		tl := len(t)
		for j := 0; j < len(nums); j ++ {

			for i := 0; i <= tl; i ++ {

			}
		}

	}
}
func main() {
	nums := []int{1,2,3,4,5,6}
	n := permute(nums)
	for _, m := range n {
		fmt.Println(m)
	}
}
