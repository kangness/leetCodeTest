package main
import (
		"fmt"
)

func reverseString(s []byte)  {
    fmt.Println(string(s))
    if len(s) <= 1 {
        return 
    }
    length := len(s)
    i := 0
    for {
        s[i], s[length - i - 1] = s[length -i -1],s[i]
        if i >= length - i - 1 {
            break
        }
        i = i + 1
    }
fmt.Println(string(s))
}
func main() {
    //s := []byte('hello')
    s := []byte{'A',' ','m','a','n',',',' ','a',' ','p','l','a','n',',',' ','a',' ','c','a','n','a','l',':',' ','P','a','n','a','m','a'}
    reverseString(s)
}

