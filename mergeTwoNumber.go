package main

import (
	"fmt"
	"math/rand"
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int)  {
	p := m + n - 1
	for {
		if m <= 0 && n <= 0 {
			break
		}
		if m - 1 >= 0 &&  (n - 1 < 0 || nums1[m-1] < nums2[n - 1]) {
			nums1[p] = nums1[m - 1]
			m = m - 1
		}else if n - 1 >= 0 && ( m - 1 < 0 || nums1[m - 1] >= nums2[n - 1]){
			nums1[p] = nums2[n - 1]
			n = n - 1
		}
		p = p - 1
	}

}

func main() {
	nums1 := make([]int, 1000)
	nums2 := make([]int, 1000)
	n1 := rand.Int() % 100
	n2 := rand.Int() % 100
	for i := 0; i < n1; i ++ {
		nums1[i] = rand.Int() % 1000
	}
	for i := 0; i < n2; i ++ {
		nums2[i] = rand.Int() % 1000
	}
	sort.Sort(sort.Reverse(sort.IntSlice(nums1)))
	sort.Sort(sort.Reverse(sort.IntSlice(nums2)))
	fmt.Println(nums1)
	fmt.Println(nums2)
	merge(nums1, n1, nums2, n2)
	fmt.Println(nums1)
	fmt.Println(n1, n2)
}

