package main
import (
	"math/rand"
	"fmt"
	"time"
)

func bubbleSort(nums []int) []int {
	if nums == nil || len(nums) <= 0 {
		return []int{}
	}
	length := len(nums)
	for i := 0; i < length ; i ++ {
		for j := i; j < length; j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	return nums
}

func main() {
	var nums []int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i <= 30; i ++ {
		nums = append(nums, rand.Intn(500))
	}
	fmt.Println(bubbleSort(nums))
}
