package main

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {
	if len(ransomNote) <= 0 && len(magazine) <= 0 {
		return true
	}
	if len(magazine) <= 0 {
		return false
	}
	r := make(map[byte]int)
	m := make(map[byte]int)
	lr := len(ransomNote)
	lm := len(magazine)
	for i := 0; i < lr; i ++ {
		r[ransomNote[i]] = r[ransomNote[i]] + 1
	}
	for i := 0; i < lm; i ++ {
		m[magazine[i]] = m[magazine[i]] + 1
	}
	for k, v := range r {
		vv, ok := m[k]
		if !ok || vv < v {
			return false
		}
	}
	return true
}

func canConstruct2(ransomNote string, magazine string) bool {
	h := make([]int, 30)
	for i := 0; i < len(magazine);i ++ {
		h[magazine[i] - 'a'] ++
	}
	for i := 0; i < len(ransomNote);i ++ {
		h[ransomNote[i] - 'a'] --
		if h[ransomNote[i] - 'a'] < 0 {
			return false
		}
	}
	return true
}
func main() {
	r := "aab"
	m := "bba"
	fmt.Println(canConstruct(r, m))
}
