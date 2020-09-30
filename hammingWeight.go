package main

func hammingWeight(num uint32) int {
	var result int
	for {
		result = int(num % 2) + result
		num = num / 2
		if num <= 0 {
			break
		}
	}
	return result

}
