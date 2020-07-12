package main
import (
	"fmt"
	"math"
)
func threeSumClosest(nums []int,target int) int{
	length := len(nums)
	var result ,total, abs, temp int
	abs = int(math.MaxInt32)
	for i := 0; i < length - 2 ; i ++ {
		for j := i + 1; j < length - 1; j ++ {
			for m := j + 1; m < length; m ++ {
				total = nums[i] + nums[j] + nums[m]
				temp = int(math.Abs(float64(target - total)))
				if temp <= abs {
					result = total
					abs = temp
				}
			}
		}
	}
	return result
}


func main() {
	nums := []int{-1,2,1,-4}
	target := 1
	fmt.Println(threeSumClosest(nums,target))
}
