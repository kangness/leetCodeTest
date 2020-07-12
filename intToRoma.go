package main

import ("fmt")

func intToRoman(num int ) string {
	value := []int{1000,900,500,400,100,90,50,40,10,9,5, 4,1}
	seps := []string{"M","CM", "D","CD", "C","XC","L","XL","X","IX","V","IV","I"}
	n := num
	var result string
	for {
		for i := 0; i < 13; i ++ {
			if n >= value[i] {
				n = n - value[i]
				result = result + seps[i]
				break
			}
		}
		if n <= 0 {
			break
		}
	}
	return result
}

func main() {
	fmt.Println(intToRoman(1994))
}
