package main

import "fmt"

func isPerfectSquare(num int) bool {
	i := 1
	for {
		num = num - i
		i = i + 2
		if num <= 0 {
			break
		}
	}
	return num == 0
}

func main() {
	num := 4
	fmt.Println(isPerfectSquare(num))
}

