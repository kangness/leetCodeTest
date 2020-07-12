package main

import (
"fmt"
"strings"
"container/list"
)
func simplifyPath(s string) string {
    l := list.New()
    l.Init()
    ss := strings.Split(s, "/")
    for _, a := range ss {
        if a == "."  || len(a) <= 0{
		continue
	}
	if a == ".." {
            if l.Back() != nil {
		l.Remove(l.Back())
	    }
		continue
	}
	l.PushBack(a)
    }
    result := ""
    for e := l.Front(); e != nil; e = e.Next() {
	result = result + "/" + e.Value.(string)
    }
    if len(result) <= 0 {
        result = "/"
    }
    return result
}
func main() {
	s := "/a//b////c/d//././/.."
//	s := "/home/"
	//s := "/../"
//	s := "/a/./b/../../c/"
//s := "/a/../../b/../c//.//"
	fmt.Println("input", s)
	fmt.Println(simplifyPath(s))

}

