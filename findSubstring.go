package main

import (
	"fmt"
)
func findSubstring(s string, words []string) []int {
	if len(words) <= 0 || len(s) <= 0 {
		return []int{}
	}
	var result []int
	worldsLen := len(words[0]) * len(words)
	slen := len(s)
	wl := len(words)
	wll := len(words[0])
	wordsMap := make(map[string]int)
	for _, w := range words {
		put(wordsMap,w)
	}
	if slen < worldsLen {
		return []int{}
	}
	for i := 0; i <= len(s) - (len(words) * wll); i ++ {
		windows := make(map[string]int)
		num := 0
		for num < wl {
			if i + num * wll >= len(s) || i + (num + 1) * wll > len(s) {
				break
			}
			w := s[i + num * wll : i + (num + 1) *wll]
			_, ok := wordsMap[w]
			if ok {
				put(windows, w)
				if windows[w] > wordsMap[w] {
					break
				}
			}else {
				break
			}
			num ++
		}
		if num == wl {
			result = append(result, i)
		}
	}
	return result
}
func put(myMap map[string]int, key string) {
	_, ok := myMap[key]
	if ok {
		myMap[key] += 1
	} else {
		myMap[key] = 1
	}
}
func main() {
	s := "wordgoodgoodgoodbestword"
	words := []string {"word","good","best","good"}
	fmt.Println(findSubstring(s,words))
}
