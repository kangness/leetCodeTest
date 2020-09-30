package main

func isUgly(num int) bool {
	if num <= 1 {
		return false
	}
	for {
		if num % 5 != 0 {
			break
		}
		num = num / 5
	}
	for {
		if num % 3 != 0 {
			break
		}
		num = num / 3
	}
	for {
		if num % 2 != 0 {
			break
		}
		num = num / 2
	}
	return num == 1
}
