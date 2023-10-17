package main

import "fmt"

func nextPermutation(nums []int) {
	n := len(nums)
	i := n - 2
	for i > 0 && nums[i] >= nums[i+1] {
		i--
	}
	if i >= 0 {
		j := n - 1
		for j > 0 && nums[i] >= nums[j] {
			j--
		}
		nums[i], nums[j] = nums[j], nums[i]
	}
	reverse(nums[i+1:])
}
func reverse(a []int) {
	for i, n := 0, len(a); i < n/2; i++ {
		a[i], a[n-1-i] = a[n-1-i], a[i]
	}
}

func main() {
	//nums := []int{1, 2, 3}
	//nums := []int{3, 2, 1}
	//nums := []int{1, 3, 2}
	nums := []int{4, 2, 0, 2, 3, 2, 0}
	nextPermutation(nums)
	fmt.Println(nums)
}
