package main

import "fmt"

func containsNearbyDuplicate(nums []int, k int) bool {

	m := make(map[int]int)
	for index, n := range nums {
		last, ok := m[n]
		if ok {
			if index - last <= k {
				return true
			}
		}
		m[n] = index
	}
	return false
}

func main() {
	//nums := []int{1,2,3,1}
	nums := []int{1,2,3,1,2,3}
	k := 1
	fmt.Println(containsNearbyDuplicate(nums,k))
}