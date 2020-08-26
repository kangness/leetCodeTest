package main

import "fmt"

func minDeletionSize(A []string) int {
	if A == nil || len(A) <= 0 {
		return 0
	}
	var a [][]byte
	for _, n := range A {
		a = append(a, []byte(n))
	}
	lA := len(A)
	la := len(A[0])
	var result int
	for j := 0; j <la; j ++ {
		for i := 1; i < lA; i ++ {
			if a[i-1][j] > a[i][j] {
				result++
				break
			}
		}
	}
	return result
}

func main() {
	A := []string{"cba", "daf", "ghi"}
	fmt.Println(minDeletionSize(A))
}
