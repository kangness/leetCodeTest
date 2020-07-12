package main

import "fmt"
func multiPly(s1, s2 string) string{
	if len(s1) <= 0 || len(s2) <= 0 {
		return ""
	}
	if s1 == "0" || s2 == "0" {
		return "0"
	}
	var result string
	l1 := len(s1) - 1
	for i := l1; i >= 0 ;i-- {
		temp := singlePly(string(s1[i]),s2)
		for j := 0; j < l1 - i; j ++ {
			temp = temp + "0"
		}
		fmt.Println(temp)
		result =  multiAdd(result, temp)
	}
	return result
}
func multiAdd(s1, s2 string) string {
	l1 := len(s1) - 1
	l2 := len(s2) - 1
	if l1 < 0 {
		return s2
	}else if l2 < 0 {
		return s1
	}
	var result string
	g := 0
	for {
		temp := 0
		if l1 < 0 && l2 < 0 {
			if g > 0 {
				result = result + string('0' + g)
			}
			break
		}
		if l1 >= 0 && l2 >= 0 {
			temp = int(s1[l1] - '0' + s2[l2] - '0') + g
			result = result + string('0' + temp % 10)
			g = temp / 10
			l1 --
			l2 --
		}else if l1 >= 0 && l2 < 0 {
			temp = int(s1[l1] - '0') + g
			result = result + string('0' + temp % 10)
			g = temp / 10
			l1 --
		}else if l1 < 0 && l2 >= 0 {
				temp = int(s2[l2] - '0') + g
				result = result + string('0' + temp % 10)
				g = temp / 10
				l2 --
		}
	}
	length := len(result)
	temp := []byte(result)
	for i := 0; i < length /2; i ++ {
		temp[i], temp[length - i - 1] = temp[length - i - 1], temp[i]
	}
	return string(temp)
}

func singlePly(s1, s2 string) string {
	length := len(s2)
	var result string
	var g int
	for i := length - 1; i >= 0;i -- {
		temp := int(s1[0] - '0') * int(s2[i] - '0') + g
		result = result + string('0' + temp % 10)
		g = temp / 10
	}
	if g > 0 {
		result = result + string('0' + g)
	}
	length = len(result)
	temp := []byte(result)
	for i := 0; i < length / 2 ; i ++ {
		temp[i],temp[length - i - 1] = temp[length - i - 1], temp[i]
	}
	return string(temp)
}
func main() {
	s1 := "123"
	s2 := "456"
	fmt.Println(multiPly(s1,s2))

}
