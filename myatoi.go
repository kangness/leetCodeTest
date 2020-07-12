package main

import (
	"fmt"
	"math"
	"strings"
)

func myAtoi(str string) int {
	fushu := false
	result := int64(0)
	num := int64(0)
	str = strings.TrimSpace(str)
	if len(str) <= 0 {
		return 0
	}
	s := []byte(str)
	if s[0] == '-' {
		s = s[1:]
		fushu = true
	}else if s[0] == '+' {
		s = s[1:]
	}
	fmt.Println(string(s))
	for _, ss := range s {
		if ss < '0' || ss > '9' {
			break
		}
		num = int64(ss) - int64('0')
		if result > math.MaxInt32 {
			if fushu {
				return math.MinInt32
			}else {
				return math.MaxInt32
			}
		}

		result = result * 10 + num
	}
	fmt.Println(result)
	if result > math.MaxInt32 {
		if fushu {
			return math.MinInt32
		}else {
			return math.MaxInt32
		}
	}else {
		if fushu {
			return int(0 - result)
		}else {
			return int(result)
		}
	}
}

func main() {
	str := "9223372036854775808"
	fmt.Println(myAtoi(str))
}
