package main

func trap (height []int) int {
	length := len(height)
	right := make([]int, length)
	left := make([]int, length)
	for i := 1; i < length; i ++ {
		if left[i -1] > height[i-1] {
			left[i] = left[i -1]
		}else {
			left[i] = height[i - 1]
		}
	}
	for i := length - 2 ; i >= 0; i -- {
		if right[i+1] > height[i+1] {
			right[i] = right[i+1]
		}else {
			right[i] = height[i+1]
		}
	}
	water := 0
	for i := 0 ; i < length; i ++ {
		level := 0
		if left[i] < right[i] {
			level = left[i]
		}else {
			level = right[i]
		}
		offset := level - height[i]
		if offset > 0 {
			water = water + offset
		}
	}
	return water
}
