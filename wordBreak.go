package main

import ("fmt")

func wordBreak(s string, wd []string) bool {
	wdm := make(map[string]int)
	for _, w := range wd {
		wdm[w] = len(w)
	}
	db := make([]bool,len(s) + 1, len(s) + 1)
	db[0] = true
	for i := 1; i <= len(s); i ++ {
		for j := 0; j < i ; j ++ {
			_, ok := wdm[s[j:i]]
			if ok && db[j] {
				db[i] = true
				break
			}
		}
	}
	return db[len(s)]
}

func main() {
	s := "aaaaaaa"
	words := []string{"aaaa", "aaa"}
	fmt.Println(wordBreak(s, words))
}
