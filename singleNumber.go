package main

import "fmt"

func singleNumber(nums []int) int {
	a := 0
	for _, n := range nums {
		a = a ^ n
	}
	return a
}

func main() {
	nums := []int{4,1,2,1,2}
	fmt.Println(singleNumber(nums))
}
