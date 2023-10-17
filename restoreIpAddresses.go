package main

import (
	"fmt"
	"strconv"
)

func stringIsOK(s string) bool {
	if len(s) <= 0 || len(s) >= 4 {
		return false
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
	if len(s) < 4 || len(s) > 12 {
		return result
	}
	for i := 3; i > 0; i-- {
		if !stringIsOK(s[:i]) {
			continue
		}
		s1 := s[i:]
		fmt.Println(s1)
		for m := 3; m > 0; m-- {
			if len(s1)-2 < m || !stringIsOK(s1[:m]) {
				continue
			}
			fmt.Println(s1, m, len(s1)-2)
			s2 := s1[m:]
			for n := 3; n > 0; n-- {
				if len(s2)-1 < n || !stringIsOK(s2[:n]) {
					continue
				}
				s3 := s2[n:]
				if !stringIsOK(s3) {
					continue
				}
				result = append(result, fmt.Sprintf("%s.%s.%s.%s", s[:i], s1[:m], s2[:n], s3))
			}
		}
	}
	return result
}

func restoreIpAddresses2(s string) []string {
	number, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if number == 0 {
		return nil
	}
	result := make(map[string]bool)
	length := len(s)
	for i := 1; i < 4; i++ {
		p1, _ := strconv.Atoi(s[:i])
		if p1 > 255 || (s[0] == '0' && (p1 != 0 || len(s[:i]) > 1)) {
			continue
		}
		for j := 1; j < 4; j++ {
			if i+j >= length {
				continue
			}
			p2, _ := strconv.Atoi(s[i : i+j])
			if p2 > 255 || (s[i] == '0' && (p2 != 0 || len(s[i:i+j]) > 1)) {
				continue
			}
			for n := 1; n < 4; n++ {
				if i+j+n >= length {
					continue
				}
				p3, _ := strconv.Atoi(s[i+j : i+j+n])
				if p3 > 255 || (s[i+j] == '0' && (p3 != 0 || len(s[i+j:i+j+n]) > 1)) {
					continue
				}
				p4, _ := strconv.Atoi(s[i+j+n:])
				if p4 > 255 || (s[i+j+n] == '0' && (p4 != 0 || len(s[i+j+n:]) > 1)) {
					continue
				}
				//key := fmt.Sprintf("%s.%s.%s.%s", s[:i], s[i:i+j], s[i+j:i+j+m], s[i+j+m:])
				key := fmt.Sprintf("%d.%d.%d.%d", p1, p2, p3, p4)
				result[key] = true
				fmt.Println(fmt.Sprintf("ip: %s.%s.%s.%s", s[:i], s[i:i+j], s[i+j:i+j+n], s[i+j+n:]))
			}
		}
	}
	var tmp []string
	for k, _ := range result {
		tmp = append(tmp, k)
	}
	return tmp
}
func main() {
	//s := "0000"
	s := "010010"
	ss := restoreIpAddresses2(s)
	fmt.Println(ss)
	fmt.Println(len(ss))
}
