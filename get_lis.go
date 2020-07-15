package main

func get_lis(a []int) int {
	length := len(a)
	if length <= 0 {
		return 0
	}
	if length == 1 {
		return 1
	}
	max := a[0]
	for i := 0; i < length; i++ {
		if max < a[i] {
			max = a[i]
		}
}
func main() {
	a := int{5, 6, 7, 1, 2, 8, 3, 10}
}
