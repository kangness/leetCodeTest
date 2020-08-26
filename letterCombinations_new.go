package main

import (
	"fmt"
)

func letterCombinations(digits string) []string {
	litter := []string{"","","abc","def","ghi","jkl","mno","pqrs","tuv","wxyz"}
	var result, tmp []string

	for index, d := range digits {
		if d <= '1' || d > '9' {
			continue
		}
		tmp = result
		result = []string{}
		offset := d - '0'
		if index == 0 {
			for _, m := range litter[offset] {
				result = append(result, string(m))
			}
			continue
		}
		for _, t := range tmp {
			for _, m := range litter[offset] {
				result = append(result, t + string(m))
			}
		}
	}
	return result
}

func main() {
	digits := "23"
	fmt.Println(letterCombinations(digits))
}
