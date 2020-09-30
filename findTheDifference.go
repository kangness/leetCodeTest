package main

import "fmt"

func findTheDifference(s, t string) byte {
	size := 26
	tmp := make([]int, 26)
	length := len(s)
	for i := 0; i < length; i ++ {
		tmp[s[i] - 'a'] ++
		tmp[t[i] - 'a'] ++
	}
	tmp[t[length] - 'a']++
	for i := 0; i < size; i ++ {
		if tmp[i] % 2 != 0 {
			return byte(i + 'a')
		}
	}
	return 'a'
}

func main() {
	s := "abcd"
	t := "abcde"
	fmt.Println(string(findTheDifference(s, t)))
}
