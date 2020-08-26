package main

import "fmt"

const badVersion int = 90
func  isBadVersion(v int) bool {
	if v >= badVersion {
		return true
	}
	return false
}

func firstBadVersion(v int) int {
	low := 0
	hight := v
	i := 0
	for  {
		i++
		mid := low + (hight - low) / 2
		fmt.Println(mid, hight,low)
		if isBadVersion(mid) {
			hight = mid
		}else {
			low = mid + 1
		}
		if low >= hight {
			break
		}
	}
	fmt.Println(i)
	return hight
}

func main() {
	v := 900
	fmt.Println(firstBadVersion(v))
}


