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
func twoSum2(nums[]int, target int) []int {
    length := len(nums)
    for i := 0; i < length; i ++ {
        for j := i +1; j <length; j++ {
            if nums[i] + nums[j] == target {
                return []int{i,j}
            }
        }
    }
    return []int{}
}

func twoSum3(nums []int, target int) []int {
    tmpMap := make(map[int][]int)
    length := len(nums)
    for i := 0 ; i < length; i ++ {
        tmpMap[nums[i]] = append(tmpMap[nums[i]],i)
    }
    for i := 0; i < length; i ++ {
        offset :=target - nums[i]
        values, ok := tmpMap[offset]
        if ok && len(values) > 0 {
            for _, v := range values {
                if v != i {
                    return []int{i,v}
                }
            }
        }
    }
    return []int{}
}
func main() {
    nums := []int{3,2,4}
    target := 6
    fmt.Println(twoSum2(nums, target))
}
