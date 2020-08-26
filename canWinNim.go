package main

import "fmt"

func canWinNim(n int) bool {
	win := make(map[int]bool)
	if n <= 3 {
		return true
	}
	win[1] = true
	win[2] = true
	win[3] = true
	for i := 5; i <= n; i ++ {
		_, ok := win[i-4]
		if ok {
			win[i] = true
		}
	}
	_, ok := win[n]
	if ok {
		return true
	}
	return false
}

func main() {
	n :=8
	fmt.Println(canWinNim(n))
}
