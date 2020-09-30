package main

func toHex(num int) string {
	hex := []string{"0","1","2","3","4","5","6","7","8","9","a","b","c","d","e","f"}
	if num == 0 {
		return hex[0]
	}
	var result string
	for {
		if len(result) >= 8 {
			break
		}
	}
}
