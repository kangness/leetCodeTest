package main

import "fmt"

func getRow(rowIndex int) []int {
	cur := 1
	var result []int
	for i:= 0; i <= rowIndex; i ++ {
		result = append(result, cur)
		cur = cur * (rowIndex - i) / (i + 1)
	}
	return result
}

func main() {
	fmt.Println(getRow(6))
}
