package main

import "fmt"

func generate(num int) [][]int {
	if num <= 0 {
		return [][]int{}
	}
	result := [][]int{}
	result = append(result, []int{1})
	for i := 1; i < num; i ++ {
		tmp := make([]int, i + 1)
		tmp[0] = 1
		for j := 1; j < i; j ++ {
			tmp[j] = result[i-1][j - 1] + result[i - 1][j]
		}
		tmp[i] = 1
		result = append(result, tmp)
	}
	return result
}

func main() {
	num := 5
	result := generate(num)
	for _, m := range result {
		fmt.Println(m)
	}
}
