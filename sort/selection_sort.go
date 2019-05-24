package main
import (
	"math/rand"
	"fmt"
	"time"
)

func selectionSort(nums []int) []int {
     if nums == nil || len(nums) <= 0 {
        return []int{}
     }
     length := len(nums)
     max := 0
     for i := 0; i < length; i ++ {
         for j := i; j < length; j++ {
		if nums[max] > nums[j] {
			max = j
		}
	} 
	nums[i],nums[max] = nums[max],nums[i]
     }
	return nums
}

func main() {
	var nums []int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i <= 30; i++ {
		nums = append(nums, rand.Intn(50))
	}
	fmt.Println(selectionSort(nums))
}
