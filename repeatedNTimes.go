package main

import "fmt"

func repeatedNTimes(A []int) int {
	tmp := make(map[int]int)
	for _, m := range A {
		_, ok := tmp[m]
		if ok {
			return m
		}
		tmp[m]++
	}
	return 0
}

func main() {
	A := []int{5,1,5,2,5,3,5,4}
	fmt.Println(repeatedNTimes(A))
}
