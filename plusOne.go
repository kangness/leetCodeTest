package main
import (
	"fmt"
)
func plusOne(digits []int) []int {
	var result []int 
	result = append(result, 0)
	result = append(result, digits...)
	length := len(result)
	result[length - 1] ++
	for i := length - 1; i >= 0; i -- {
		if result[i] >= 10 {
			result[i] = result[i] % 10
			if i > 0 {
				result[i - 1] = result[i -1] + 1
			}
		}
	}
	if result[0] > 0 {
		return result
	}
	return result[1:]
}

func main() {
	nums := []int{1,2,3,4,5,6,7,8,9}
	fmt.Println(plusOne(nums))
}
