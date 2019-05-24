package main
import (
		"fmt"
)

func maxArea(nums []int) int {
    length := len(nums)
    var area, max int
    for l := length -1; l > 0; l -- {
        for i := 0;i < length - l; i++ {
		if nums[i] < nums[i+l] {
                    area = nums[i] * l    
		}else {
                    area = nums[i + l] * l    
		}
		if max <= area {
			max = area
		}
	}
    }
	    return max
}
func main() {
    nums1 := []int{1,8,6,2,5,4,8,3,7}
    fmt.Println(maxArea(nums1))
}
