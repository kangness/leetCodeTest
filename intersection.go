package main

import "fmt"

func intersection(nums1, nums2 []int) []int {
	if nums1 == nil || nums2 == nil || len(nums1) <= 0 || len(nums2) <= 0 {
		return nil
	}
	n1 := make(map[int]bool)
	n2 := make(map[int]bool)
	for _, n := range nums1 {
		n1[n] = true
	}
	for _, n := range nums2 {
		n2[n] = true
	}
	var result []int
	for k, _ := range  n1 {
		_, ok := n2[k]
		if ok {
			result = append(result, k)
		}
	}
	return result
}

func main () {
	nums1 := []int{1,2,2,1}
	nums2 := []int{2,2}
	fmt.Println(intersection(nums1, nums2))
}
