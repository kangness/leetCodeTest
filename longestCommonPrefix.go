package main
import (
	"fmt"
)

func longestCommonPrefix(strs []string) string {
	length := len(strs)
	if length <= 0 {
		return ""
	}
	var result string
	j := 0
	for {
		isTrue := true
		value := strs[0][j]
		for i := 0; i < length; i ++ {
			if j + 1 > len(strs[i])  {
				return result
			}
			if strs[i][j] != value {
				isTrue = false
				break
			}
		}
		if isTrue {
			result = strs[0][:j+1]
		}else {
			break
		}
		j ++
	}
	return result
}

func main() {
	strs := []string{"flower","flow","flight"}
	fmt.Println(longestCommonPrefix(strs))
}
