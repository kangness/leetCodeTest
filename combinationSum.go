package main
import (
	"fmt"
)

func combinationSum(cand []int, target int) [][] int {
	if len(cand) <= 0 {
		return [][]int{}
	}
	if len(cand) == 1 {
		if cand[0] == target {
			return [][]int{cand}
		}
		return [][]int{}
	}
	for i := 0; i < len(cand) - 1 ; i ++ {
		for j := i; j < len(cand) ; j ++ {
			if cand[i] > cand[j] {
				cand[i], cand[j] = cand[j], cand[i]
			}
		}
	}
	temp := []int{}
	return combin(cand,temp, target, 0)
}

func combin(cand, temp []int, target , n int) [][]int {
	var result [][]int
	if target == 0 {
		result = append(result, temp)
	}
	if target < cand[0] {
		return result
	}
	for i := n;i < len(cand) ; i ++ {
		if cand[i] > target {
			break
		}
		tmp := []int{}
		tmp = append(tmp, temp...)
		tmp = append(tmp, cand[i])
		tt := combin(cand,tmp,target - cand[i],i)
		if tt != nil && len(tt) > 0 {
			result = append(result, tt...)
		}
	}
	return result
}


func main() {
	cand := []int{2,3,6,7}
	target := 7
	result := combinationSum(cand,target)
	for _, m := range result {
		fmt.Println(m)
	}
}
