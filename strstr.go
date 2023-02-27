package main

import "fmt"

func strStr(haystack string, needle string) int {
	hlen := len(haystack)
	nlen := len(needle)
	if nlen > hlen {
		return -1
	}
	for i := 0; i < hlen; i ++ {
		right := 0
		for j := 0; j < nlen; j ++ {
			if i + j >= hlen {
				break
			}
			if haystack[i+j] != needle[j] {
				break
			}
			right ++
		}
		if right == nlen {
			return i
		}
	}
	return -1
}

func main() {
	haystack := "abc"
	needle := "c"
	fmt.Println(strStr(haystack,needle))
}
