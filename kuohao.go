package main

import ("fmt")

func Perm(str []byte, begin, end int) []string {
	var result []string
	if begin == end {
		result = append(result, string(str))
		return result
	}
	for i := begin; i <= end; i ++ {
		if i != begin {
			str[begin],str[i] = str[i], str[begin]
		}
		result = append(result,Perm(str, begin + 1, end)...)
		if i != begin {
			str[begin],str[i] = str[i], str[begin]
		}
	}
	return result
}

func main() {
	str := []byte("()")
	result := Perm(str,0,5)
	for _, m := range result {
		fmt.Println(string(m))
	}
}
