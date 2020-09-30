package main

import (
	"container/list"
	"fmt"
)

func buildArray(target []int, n int) []string {
	tmp := list.New()
	for i := 1; i <= n; i ++ {
		tmp.PushBack(i)
	}
	var result []string
	for _, v := range target {
		tv := tmp.Front().Value.(int)
		if tv == v {
			result = append(result, "Push")
		}else {
			for {
				tt := tmp.Front().Value.(int)
				if tt != v {
					result = append(result, "Push")
					result = append(result, "Pop")
					tmp.Remove(tmp.Front())
				}else {
					result  = append(result, "Push")
					break
				}
			}
		}
		tmp.Remove(tmp.Front())
	}
	return result
}

func main() {
	target := []int{2,3,4}
	n := 4
	fmt.Println(buildArray(target, n))
}
