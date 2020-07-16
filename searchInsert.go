package main
import (
    "fmt"
)
func searchInsert(nums []int, target int) int {
    if nums == nil || len(nums) <= 0 || nums[0] >= target {
        return 0
    }
    length := len(nums)
    if nums[length - 1] < target {
        return length
    }
    for i := 1; i < length; i ++ {
        if target == nums[i] {
            return i
        }
        if target > nums[i - 1] && target < nums[i] {
            return i
        } 
    }
    return length
}
func main() {
    //nums := []int{1,3,5,6}
    nums := []int{1}
    fmt.Println(searchInsert(nums, 1))
}
