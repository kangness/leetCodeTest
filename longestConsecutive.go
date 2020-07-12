package main

import (
	"fmt"
)

func longestConsecutive(nums []int) int {
	temp := make(map[int]int)
	for _, n := range nums {
		temp[n] = temp[n] + 1
	}
	max := 0
	for k, _ := range temp {
		count := 1
		i := 1
		for {
			_, ok := temp[k + i]
			if ok {
				count = count + 1
				delete(temp,k + i)
				i = i + 1
			}else {
				break
			}
		}
		i = 1
		for {
			_, ok := temp[k - i]
			if ok {
				count = count + 1
				delete(temp, k - i)
				i ++
			}else {
					break
			}
		}
		if max < count {
			max = count
		}
	}
	return max
}

func main() {
	nums := []int{100, 4, 200, 1, 3, 2}
	fmt.Println(longestConsecutive(nums))
}
