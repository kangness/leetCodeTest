package main

import "fmt"

func numDecodings(s string) int{
	dp := make([]int, len(s) + 1)
	length := len(s)
	if length <= 0 || (length == 1 && s[0] == '0') {
		return 0
	}
	if length <= 1 {
		return 1
	}
	dp[0] = 1
	for i := 0; i < length; i ++ {
		if s[i] == '0' {
			dp[i + 1] = 0
		}else {
			dp[i + 1] = dp[i]
		}
		if i > 0 && (s[i - 1] == '1' || (s[i -1] == '2' && s[i] <= '6')) {
			dp[i]

		}
	}
}

func main() {
	num := "222222"
	fmt.Println(numDecodings(num))
}
