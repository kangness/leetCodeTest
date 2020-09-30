package main

import "fmt"

func convertToTitle(n int) string {
	var result string
	for {
		n --
		result =  string(n % 26 + 'A') + result
		n = n  / 26
		if n <= 0 {
			break
		}
	}
	return result
}

func main() {
	n := 28
	fmt.Println(convertToTitle(n))
}
