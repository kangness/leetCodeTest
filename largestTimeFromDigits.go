package main

import (
	"fmt"
	"sort"
)

func largestTimeFromDigits(A []int) string {
	sort.Ints(A)

	for i := 3; i >= 0; i -- {
		if A[i] > 2 {
			continue
		}
		for j := 3 ; j >= 0; j -- {
			if j == i || (A[i] == 2 && A[j] > 3) {
				continue
			}
			for n := 3; n >= 0; n -- {
				if n == j || n == i || A[n] > 5 {
					continue
				}
				for m := 3; m >= 0 ; m -- {
					if m == i || m == j || m == n {
						continue
					}
					return fmt.Sprintf("%d%d:%d%d",A[i],A[j],A[n],A[m])
				}
			}

		}
	}
	return ""
}

func main() {
	//A := []int {1,2,3,4}
	A := []int {5,5,5,5}
	fmt.Println(largestTimeFromDigits(A))
}
