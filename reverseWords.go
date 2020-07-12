package main

import (
    "strings"
    "fmt"
)
func reverseWords(s string) string {
    ss := strings.Split(s," ")
    length := len(ss)
    result := ""
    for i := length - 1; i >= 0; i-- {
        if len(ss[i]) <= 0 {
            continue
        }
        result = result + " " + ss[i]
    }
    if result[0] == ' ' {
        result = result[1:]
    }
    return result
}
func main() {
	s := "the sky is blue"
	fmt.Println(reverseWords(s), "&&&")
}
