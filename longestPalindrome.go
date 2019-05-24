package main
import (
		"fmt"
)

func longestPalindrome(s string) string {
    ss := []byte(s)
    length := len(ss)
    for l := length ; l > 0; l-- {
        for j := 0; j < length - l; j ++ {
            m := j
            n := j + l - 1
            for {
                if ss[m] != ss[n] {
                    break
                }
                m = m + 1
                n = n - 1
                if n <= m {
                    return string(ss[j:j+l])
                }
            }
        }
    }
    return ""

}

func main() {
	s := "babad" 
   fmt.Println(longestPalindrome(s))   
}
