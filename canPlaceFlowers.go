package main

import "fmt"

func canPlaceFlowers(flowerbed []int, n int) bool {
	if len(flowerbed) == 1 && flowerbed[0] == 0 && n <= 1 {
		return true
	}
	for i := 0; i < len(flowerbed); i ++ {
		if flowerbed[i] == 0 {
			if i != 0 && i < len(flowerbed) - 1 {
				if flowerbed[i - 1] == 0 && flowerbed[i + 1] == 0 {
					n --
					flowerbed[i] = 1
				}
			}
			if i == 0 {
				if i + 1 < len(flowerbed) &&  flowerbed[i + 1] == 0 {
					n --
					flowerbed[i] = 1
				}
			}
			if i == len(flowerbed) - 1 {
				if  i - 1 >= 0 && flowerbed[i - 1] == 0 {
					n --
					flowerbed[i] = 1
				}
			}
		}
		if n <= 0 {
			return true
		}

	}
	return false
}

func main() {
	flowerbed := []int {1,0,0,0,1}
	n := 2
	fmt.Println(canPlaceFlowers(flowerbed, n))
}
