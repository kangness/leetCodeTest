package main
import (
   "fmt"
)
func checkInclusion(s1, s2 string) bool {
        if len(s1) > len(s2) {
		return false
	}
        size := 512
        a1 := make([]int, size)
        a2 := make([]int, size)
	length2 := len(s2)
	length1 := len(s1)
        for i := 0; i < size; i ++ {
            a1[i] = 0
            a2[i] = 0
        }
        for i := 0; i < len(s1) ;i ++ {
		a1[int(s1[i])] ++
        }
        for i := 0; i <= length2 - length1; i ++ {
                for j := 0; j < size; j ++ {
			a2[j] = 0
		} 
		for j := 0; j < length1; j ++ {
			a2[int(s2[i+j])] ++
		}
		isOk := true
		for j := 0; j < size; j ++ {
			if a1[j] != a2[j] {
				isOk = false
				break
			}
		}
		if isOk == true {
			return true
		}
	}
	return false
}
func  main() {
    s1 := "ab"
    s2 := "eidbaooo"
    fmt.Println(checkInclusion(s1, s2)) 

}
