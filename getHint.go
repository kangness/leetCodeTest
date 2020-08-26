package main

import "fmt"

func getHint(secret string, guess string) string {
	ss := []byte(secret)
	gg := []byte(guess)
	A := 0
	B := 0
	lengthG := len(gg)
	for i, m := range ss {
		if i < lengthG && gg[i] == m {
			A ++
		}
	}
	for i, m := range gg {
		for j , n := range ss {
			if m == n && i != j {
				B ++
				break
			}
		}
	}
	return fmt.Sprintf("%dA%dB", A,B)
}

func main() {
	secret := "1123"
	guess := "0111"
	fmt.Println(getHint(secret,guess))
}
