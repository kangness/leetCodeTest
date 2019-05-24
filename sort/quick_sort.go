package main

import (
	"fmt"
	"time"
	"math/rand"
)

func quickSort(nums []int) []int {
	if nums == nil || len(nums) <= 1 {
		return nums
	}
	head, tail := 0, len(nums) - 1
	mid, i := nums[0], 1
	for head < tail {
		if nums[i] > mid {
			nums[i], nums[tail] = nums[tail], nums[i]
			tail --
		} else {
			nums[i] , nums[head] = nums[head], nums[i]
			head ++
			i ++
		}
	}
	nums[head] = mid
	quickSort(nums[:head])
	quickSort(nums[head+1:])
	return nums
}

func main() {
	rand.Seed(time.Now().UnixNano())
	var nums []int
	for i := 0; i < 30; i++ {
		nums = append(nums, rand.Intn(1000))
	}
	fmt.Println(quickSort(nums))
}

