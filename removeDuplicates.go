package main
import (
     "fmt"
)
func removeDuplicates(nums []int) int {
    if len(nums) < 2 {
        return len(nums)
    }
    length := len(nums)
    j := 0
    for i := 1; i < length; i ++ {
        if nums[j] != nums[i] {
            j = j + 1
            nums[j] = nums[i]
        }
    }
    return j + 1
}

func main() {
    //nums := []int{0,0,1,1,1,2,2,3,3,4}
    //nums := []int{1,1,2}
    nums := []int{1,2,3,4,5,6,7,8,8}
    fmt.Println(nums)
    fmt.Println(removeDuplicates(nums))
    fmt.Println(nums)
}
