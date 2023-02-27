package main

import "fmt"

func addDigits(num int) int {
	result := 0
	for {
		result = result +  num % 10
		num = num / 10
		if num < 10 {
			result = result + num
			num = result
			if result >= 10 {
				result = 0
			}
		}
		if num < 10 && result < 10 {
			break
		}
	}
	return result
}

func addDigits2(num int) int {
	return (num - 1) % 9 + 1;
}
func main() {
	num := 38
	fmt.Println(addDigits(num))
}
