package main

import (
	"fmt"
)

func convertZ(s string, n int) string {
	ss := []byte(s)
	var temp [][]byte
	l := len(ss)
	ssss := "                                                                          "
	for i := 0; i < n; i++ {
		temp = append(temp, []byte(ssss))
	}
	offset := 0
	j, i := -1, 0
	down := true
	for {
		fmt.Println(i, j, offset)
		if offset >= l {
			break
		}
		if down {
			j++
		} else {
			j--
		}
		temp[j][i] = ss[offset]
		if j == 0 || j == n-1 {
			i = i + 1
			if j == 0 {
				down = true
			} else {
				down = false
			}
		}
		offset++
	}
	fmt.Println("")
	for i := 0; i < n; i++ {
		fmt.Println(string(temp[i]))
	}
	return ""
}

func main() {
	s := "0123456789"
	fmt.Println(convertZ(s, 3))
}
