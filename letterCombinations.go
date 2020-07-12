package main
import (
	"fmt"
	"container/list"
)

func letterCombinations(digits string) []string {
	litter := []string{"","","abc","def","ghi","jkl","mno","pqrs","tuv","wxyz"}
	l := list.New()
	l.PushFront("")
	length := len(digits)
	result := []string{}
	for i := 0 ;i < length; i ++ {
		num := int(digits[i] - '0')
		if num < 2 || num > 9 {
			continue
		}
		for {
			if l.Front() == nil || len(l.Front().Value.(string)) != i {
				break
			}
			ll := len(litter[num])
			temp := l.Remove(l.Front()).(string)
			for j := 0; j < ll; j ++ {
				l.PushBack(temp + string(litter[num][j]))
			}
		}
	}
	for e := l.Front(); e != nil ; e = e.Next() {
		fmt.Println(e.Value.(string))
		result = append(result, e.Value.(string))
	}
	return result
}

func letterCombinations2(digits string) []string {
	if len(digits)==0{
		return nil
	}
	m := map[byte][]string{
		'2': {"a", "b", "c"},
		'3': {"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': {"j", "k", "l"},
		'6': {"m", "n", "o"},
		'7': {"p", "q", "r", "s"},
		'8': {"t", "u", "v"},
		'9': {"w", "x", "y", "z"},
	}
	s := m[digits[0]]
	for i := 1; i < len(digits); i++ {
		temp := make([]string, 0)
		for j := 0; j < len(s); j++ {
			for k := 0; k < len(m[digits[i]]); k++ {
				temp = append(temp, s[j] + m[digits[i]][k])
			}
		}
		fmt.Println("--------", temp)
		s = temp
	}
	return s
}
func main() {
	digits := "2345"
	fmt.Println(letterCombinations2(digits))
}
