package main

import "fmt"

func addStrings(num1, num2 string) string {
	length1 := len(num1)
	length2 := len(num2)
	length := 0
	if length1 > length2 {
		length = length1
	} else {
		length = length2
	}
	var l1, l2,l3, over int
	var result string
	for i := 1; i <= length;i ++ {
		l1 = 0
		l2 = 0
		l3 = 0
		if length1 - i >= 0 {
			l1 = int(num1[length1 - i] - '0')
		}
		if length2 - i >= 0 {
			l2 = int(num2[length2 - i ] - '0')
		}
		l3 = l1 + l2 + over
		over = l3 / 10
		l3 = l3 % 10
		result = string(byte(l3 + '0'))  + result
	}
	if over > 0 {
		result = string(byte('0' + over )) + result
	}
	return result
}

func main() {
	num1 := "5000"
	num2 := "5000"
	fmt.Println(addStrings(num1, num2))
}
