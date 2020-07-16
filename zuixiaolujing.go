package main

import (
	"fmt"
)

func minRoad(road [][]int) int {
	dp := make([][]int, len(road)+1)
	for i := 0; i <= len(road); i++ {
		dp[i] = make([]int, len(road[0])+1)
	}
	dp[0][0] = road[0][0]
	t1 := 999999
	t2 := 999999
	for i := 0; i < len(road); i++ {
		for j := 0; j < len(road[0]); j++ {
			if i == 0 && j == 0 {
				continue
			}
			if j-1 >= 0 {
				t1 = dp[i][j-1] + road[i][j]
			}
			if i-1 >= 0 {
				t2 = dp[i-1][j] + road[i][j]
			}
			if t1 < t2 {
				dp[i][j] = t1
			} else {
				dp[i][j] = t2
			}
		}
	}
	fmt.Println(dp)
	return dp[len(road)-1][len(road[0])-1]
}
func main() {
	var road [][]int
	road = append(road, []int{1, 3, 1})
	road = append(road, []int{1, 5, 1})
	road = append(road, []int{4, 2, 1})
	fmt.Println(road)
	fmt.Println(minRoad(road))
}
