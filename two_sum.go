package main
import (
    "fmt"
)
func twoSum(nums[]int, target int) []int {
    length := len(nums)
    for i := 0; i <= length - 1; i++ {
        for j := i; j <= length - 1; j++ {
            if (nums[i] + nums[j] == target) {
                return []int{i,j}
            }
        }
    }
    return []int{}
}

func
func main() {
    nums := []int{3,2,4}
    target := 6
    fmt.Println(twoSum(nums, target))
}
