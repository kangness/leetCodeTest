package main
import ("fmt")

func fourSum(nums []int, target int) [][]int {
	length := len(nums)
	var result [][]int
	for i := 0; i < length - 1; i ++ {
		for j := i + 1; j < length; j ++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}
	quchong := make(map[string]bool)
	for i := 0; i < length - 3; i ++ {
		for j := i + 1; j < length - 2; j ++ {
			m := j + 1
			n := length - 1
			for {
				if m >= n {
					break
				}
				if nums[i]+nums[j]+nums[m]+nums[n] < target {
					m++
				} else if nums[i]+nums[j]+nums[m]+nums[n] > target {
					n--
				} else {
					key := fmt.Sprintf("%d_%d_%d_%d", nums[i], nums[j], nums[m], nums[n])
					_, ok := quchong[key]
					if !ok {
						result = append(result, []int{nums[i], nums[j], nums[m], nums[n]})
						quchong[key] = true
					}
					m ++
					n --
				}
			}
		}

	}
	return result
}

func main() {
	nums := []int{-3,-2,-1,0,0,1,2,3}
	target := 0
	result := fourSum(nums, target)
	for _, m := range result {
		fmt.Println(m)
	}
}
