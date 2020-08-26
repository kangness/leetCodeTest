package main

import "fmt"

func sqrt(x int) int {
	if x <= 1 {
		return 1
	}
	var max, min, m  int
	max = x
	for {
		if max - min <= 1 {
			break
		}
		m = (max + min) / 2
		if x / m > max {
			min = m
		}else {
			max = m
		}
	}
	return min
}

func main() {
	x := 9
	fmt.Println(sqrt(x))
}
