package main
import (
		"fmt"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
    length1 := len(nums1)
    length2 := len(nums2)
    var l1,l2 int
    var nums3 []int
    for {
        if l1 < length1 && l2 < length2 {
            if nums1[l1] <= nums2[l2] {
                nums3 = append(nums3, nums1[l1])
                l1 = l1 + 1
            }else {
                nums3 = append(nums3, nums2[l2])
                l2 = l2 + 1
            }
        }
        if l1 >= length1 && l2 < length2 {
            nums3 = append(nums3, nums2[l2])
            l2 = l2 + 1
        }
        if l2 >= length2 && l1 < length1 {
            nums3 = append(nums3, nums1[l1])
            l1 = l1 + 1
        }
        if l2 >= length2 && l1 >= length1 {
            break
        }
    }
    length := len(nums3)
    if length % 2 == 0 {
        return float64(float64(nums3[length/2 -1] + nums3[length/2])/2)
    }else {
        return float64(nums3[int(length/2)])
    }
}
func main() {
    nums1 := []int{1,2,3}
    nums2 := []int{3,4}
    fmt.Println(findMedianSortedArrays(nums1,nums2))
}
