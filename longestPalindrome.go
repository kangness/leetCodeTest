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

/*
给定一个包含大写字母和小写字母的字符串 s ，返回 通过这些字母构造成的 最长的回文串 。

在构造过程中，请注意 区分大小写 。比如 "Aa" 不能当做一个回文字符串。
示例 1:

输入:s = "abccccdd"
输出:7
解释:
我们可以构造的最长的回文串是"dccaccd", 它的长度是 7。
 */
func longestPalindrome2(s string) int {
    tmpMap := make(map[byte]int)
    ss := []byte(s)
    for _, s := range ss {
        tmpMap[s]++
    }
    if len(tmpMap) <= 1 {
        return len(s)
    }
    var hasJi bool
    hasJi = false
    totalLength := 0
    for k , v := range tmpMap {
        fmt.Println(string(k),v)
        totalLength = totalLength + v
        if v % 2 == 1 {
            if hasJi {
                totalLength = totalLength - 1
            }else {
                hasJi = true
            }
        }
    }
    return totalLength
}
func main() {
	s := "tattarrattat"
   fmt.Println(longestPalindrome(s))
    fmt.Println(longestPalindrome2(s))
}
/*
7
 */
