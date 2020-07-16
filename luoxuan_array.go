package main

import (
	"fmt"
)

func luoxuan(nums [][]int) []int {
	var result []int
	total := 1
	ll := len(nums) - 1
	lr := len(nums[0]) - 1
	total = total + len(nums)*len(nums[0])
	var i, j, m, n, st int
	fmt.Println("total:%d", total)
	for t := 1; t < total; t++ {
		result = append(result, nums[i][j])
		fmt.Println(fmt.Sprintf("total :%d st:%d ll:%d lr:%d i:%d j:%d m:%d n:%d num:%d", t, st, ll, lr, i, j, n, m, nums[i][j]))
		if st%4 == 0 {
			if j < lr {
				j = j + 1
			} else {
				st = st + 1
				i = i + 1
			}
			continue
		}
		if st%4 == 1 {
			if i < ll {
				i = i + 1
			} else {
				st = st + 1
				j = j - 1
			}
			continue
		}

		if st%4 == 2 {
			if j > m {
				j = j - 1
			} else {
				st = st + 1
				i = i - 1
			}
			continue
		}
		if st%4 == 3 {
			if i > n+1 {
				i = i - 1
			} else {
				st = st + 1
				n = n + 1
				m = m + 1
				ll = ll - 1
				lr = lr - 1
				i = n
				j = m
			}
			continue
		}

	}
	return result
}
func main() {
	var nums [][]int
	total := 16
	length := 4
	for i := 0; i < total/length; i++ {
		nums = append(nums, make([]int, length))
	}
	for i := 0; i < total; i++ {
		nums[i/length][i%length] = i + 1
	}
	fmt.Println(nums)
	fmt.Println(luoxuan(nums))
}
