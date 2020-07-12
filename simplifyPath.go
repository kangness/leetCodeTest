package main

import (
		"fmt"
	"contain/list"
)
func simplifyPath(s string) string {
    for {
	if strings.Contains(s, "//") {
		s = strings.Replace(s, "//", "/", -1)
	}else {
		break
	}
    }
    for {
        if strings.Contains(s, "/./") {
		s = strings.Replace(s, "/./", "/", -1)
	}else {
		break
	}
    }
    fmt.Println("begin", s)
    for {
         fmt.Println("start", s)
         index := strings.Index(s, "/..")
         if index < 0 {
             break
         }
         if index == 0 {
              s = s[3:]
	      continue
         }
         i := index - 1
         for i = index - 1; i >= 0; i-- {
             if s[i] == '/' {
                 break
	     }
         }
         if i == 0 {
             s = s[index+3:]
	     continue
         }
	 fmt.Println("debug", s[:i], s[index + 3:] )
         s = s[:i] + s[index + 3:]
    }
    length := len(s)
    if length > 1 && s[length -1] == '/' {
        s = s[:length - 1]
    }
    return s
}
func main() {
	//s := "/a//b////c/d//././/.."
//	s := "/home/"
	s := "/../"
//	s := "/a/./b/../../c/"
//s := "/a/../../b/../c//.//"
	fmt.Println("input", s)
	fmt.Println(simplifyPath(s))
}
