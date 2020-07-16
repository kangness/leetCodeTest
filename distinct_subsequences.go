package main

import (
	"fmt"
)

func distinctSubsequences(s, t string) int {
	var dp [][]int
	ss := []byte(s)
	tt := []byte(t)
	sl := len(ss)
	st := len(tt)
	dp = make([][]int, sl+1)
	for i := 0; i < sl; i++ {
		dp[i] = make([]int, st+1)
		dp[i][0] = 1
	}
	if sl < st || (sl == st && s != t) {
		return 0
	}
	dp[0][0] = 1
	for j := 1; j < st; j++ {
		dp[0][j] = 0
	}
	for i := 1; i < sl; i++ {
		for j := 1; j < st; j++ {
			if ss[i] == ss[j] {
				dp[i][j] = dp[i-1][j-1] + dp[i-1][j]
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	fmt.Println(dp)
	fmt.Println(6 & 1)
	return dp[sl-1][st-1]
}
func main() {
	s := "arabbbit"
	t := "rabbit"
	fmt.Println(distinctSubsequences(s, t))
}
