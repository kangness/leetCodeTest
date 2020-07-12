package main

import (
		"fmt"
		"strconv"
       )

func stringIsOK(s string) bool {
	if len(s) <= 0 || len(s) >= 4 {
		return false;
	}
	num, err := strconv.Atoi(s)
		if err != nil {
			fmt.Println(err)
				return false
		}
	if num > 255 {
		return false
	}
	if num < 10 && len(s) > 1 {
		return false
	}
	if num < 100 && len(s) > 2 {
		return false
	}
	return true
}
func restoreIpAddresses(s string) []string {
	var result []string
		if len(s) < 4 || len(s)  > 12 {
			return result
		}
	for i := 3; i > 0; i -- {
		if !stringIsOK(s[:i]) {
			continue
		}
s1 := s[i:]
	    fmt.Println(s1)
	    for m := 3; m > 0; m -- {
		    if len(s1) - 2  < m || !stringIsOK(s1[:m]) {
			    continue
		    }
		    fmt.Println(s1, m, len(s1) - 2)
			    s2 := s1[m:]
			    for n := 3; n > 0; n -- {
				    if len(s2) - 1 < n || !stringIsOK(s2[:n]) {
					    continue
				    }
s3 := s2[n:]
	    if !stringIsOK(s3) {
		    continue
	    }
    result = append(result, fmt.Sprintf("%s.%s.%s.%s",s[:i], s1[:m],s2[:n],s3))
			    }
	    }
	}
	return result
}

func main() {
s := "0000"
	   fmt.Println(restoreIpAddresses(s))
}
