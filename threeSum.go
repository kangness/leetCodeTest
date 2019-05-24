package main
import (
		"fmt"
)
func threeSum(nums []int) [][]int {
    length := len(nums)
    number := 0
    for i := 0; i < length; i ++ {
        for j := i; j < length; j++ {
            if nums[i] > nums[j] {
                nums[i], nums[j] = nums[j], nums[i]
            }
        }
    }
	fmt.Println(nums)
    tempMap := make(map[string]bool)
    var result [][]int
    var m, n int
    for i := 0; i < length - 1; i++ {
	m = i + 1
	n = length - 1
        if nums[i] > 0 {
		break
	}
	for {
		if m >= n {
			break
		}
		number = nums[i] + nums[m] + nums[n]
		if number > 0 {
			n = n - 1
		}
		if number < 0 {
			m = m + 1
		}
		if number == 0 {
			key := getMapKey(nums[i],nums[m],nums[n])
			_, ok := tempMap[key]
			if !ok {
				tempMap[key] = true
				result = append(result, []int{nums[i],nums[m],nums[n]})
			}
			n = n - 1
		}
	}
    }
    fmt.Println(result)
    return result
}
func getMapKey(a,b,c int) string {
    
    if(a>b) {
        a,b = b, a
    }
    if(b>c) {
        b,c = c, b
    };
    if(a>b) {
        a,b = b, a
    }
    return fmt.Sprintf("%d_%d_%d",a,b,c)
}
func main() {
   num1 := []int{-1, 0, 1, 2, -1, -4}
   threeSum(num1)   
}
