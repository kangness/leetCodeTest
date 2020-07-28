package main
import (
     "fmt"
)
func removeDuplicates(nums []int) int {
    if nums == nil || len(nums) <= 0 {
        return 0
    }
    if len(nums) <= 1 {
        return 1
    }
    var tmp int
    tail := len(nums) - 1
    head := 0
    for i := 0; i < tail; i ++ {
        if nums[i] == nums[i + 1] {
            tmp = nums[i]
            nums[i] = nums[head]
            nums[head] = tmp
            fmt.Println(fmt.Sprintf("tail:%d head:%d i:%d", tail, head, i), nums)
	    head++
	}
    }
    nums = nums[head:]
    fmt.Println(nums)
    return len(nums)
}

func main() {
    //nums := []int{0,0,1,1,1,2,2,3,3,4}
    nums := []int{1,1,2}
    fmt.Println(nums)
    fmt.Println(removeDuplicates(nums))
    fmt.Println(nums)
}
