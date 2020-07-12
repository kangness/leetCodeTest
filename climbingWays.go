package main

import (
	"fmt"
)

func getClibingWays(s int) int {
	tmp := 0
	if s < 1 {
		return 0
	}
	if s == 1 {
		return 1
	}
	if s == 2 {
		return 2
	}
	a := 1
	b := 2
	for i := 3; i <= s; i++ {
		tmp = a + b
		a = b
		b = tmp
	}
	return tmp
}

func step(s int) int {
	result := 0
	if s <= 1 {
		return 1
	}
	if s == 2 {
		return 2
	}
	if s > 2 {
		result = step(s-1) + step(s-2)
	}
	return result
}
func main() {
	for i := 1; i < 10; i++ {
		fmt.Println(i, step(i))
		fmt.Println(i, getClibingWays(i))
	}
}
