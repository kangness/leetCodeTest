package main

import "fmt"

func fizzBuzz(n int ) []string {
	var result []string
	for i := 1; i <= n; i ++ {
		if i % 15 == 0 {
			result = append(result, "FizzBuzz")
		}else if i % 3 == 0 {
			result = append(result, "Fizz")
		}else if i % 5 == 0 {
			result = append(result, "Buzz")
		}else {
			result = append(result, fmt.Sprintf("%d", i))
		}
	}
	return result
}

func main () {
	n := 15
	fmt.Println(fizzBuzz(n))
}
