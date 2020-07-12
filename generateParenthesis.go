package main
import (
"fmt"
)
func generateParenthesis(n int) []string {
	return gp("",0,0,n)
}

func gp(s string, l, r, n int) []string {
	var result []string
	if l > n || r > n || r > l {
		return result
	}
	if l == n && r == n {
		result = append(result, s)
		return result
	}
	res := gp(s + "(", l+1, r, n)
	if res != nil && len(res) > 0 {
		result = append(result, res...)
	} 
	res = gp(s + ")", l, r + 1, n)
	if res != nil && len(res) > 0 {
		result = append(result, res...)
	} 
	return result
}
func main () {
	n := 3
	fmt.Println(generateParenthesis(n))
}
