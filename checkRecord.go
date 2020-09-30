package main

import "fmt"

func checkRecord(s string) bool {
	var a int
	for index, m := range s {
		if m == 'A' {
			a ++
		}
		if index >= 2 && s[index] == 'L' && s[index-1] == 'L' && s[index-2] == 'L' {
			return false
		}
		if a > 1 {
			return false
		}
	}
	return true
}

func main() {
	s := "PPALLL"
	fmt.Println(checkRecord(s))
}

