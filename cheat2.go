package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i <= 10; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Printf("loop i is - %d\n", i)
		}(i)
	}
	wg.Wait()
	fmt.Println("Hello, Welcome to Go")
}
