package main

import (
	"fmt"
	)

func romanToInt(r string) int {
	d := map[string]int{"M":1000,"D":500,"C":100,"L":50,"X":10,"V":5,"I":1}
	old := "M"
	s := []byte(r)
	result := 0
	for _, m := range s {
		result = d[string(m)] + result
		if d[string(m)] > d[old] {
			result = result - d[old] - d[old]
		}
		old = string(m)
	}
	return result
}
func main() {
	r := "MCMXCIV"
	fmt.Println(romanToInt(r))
}
