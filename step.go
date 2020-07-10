package main

import (
	"fmt"
)

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
	}
}
