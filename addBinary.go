package main
import ("fmt")

func addBinary(a, b string) string {
	la := len(a) - 1
	lb := len(b) - 1
	last := 0
	temp := 0
	var result string
	for {
		if la < 0 && lb < 0 {
			if last > 0 {
				result = result + "1"
			}
			break
		}
		if la >= 0 && lb >= 0 {
			temp = int(a[la]  - '0') + int(b[lb] - '0') + last
			last = temp / 2
			temp = temp % 2
			la --
			lb --
		}else if la >= 0 && lb < 0 {
				temp = int(a[la] - '0') + last
				last = temp / 2
				temp = temp % 2
				la --
		}else if la < 0 && lb >= 0 {
			temp = int(b[la] - '0') + last
			last = temp / 2
			temp = temp % 2
			lb --
		}
		if temp == 0 {
			result = result + "0"
		}else {
			result = result + "1"
		}
	}
	length := len(result)
	tmp := []byte(result)
	for i := 0; i < length /2 ; i ++ {
		tmp[i], tmp[length - i - 1] = tmp[length - i -1], tmp[i]
	}
	return string(tmp)
}

func main () {
	a := "1010"
	b := "1011"
	fmt.Println(addBinary(a,b))
}
