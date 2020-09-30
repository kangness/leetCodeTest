package main

import (
	"fmt"
	"math"
	"sync"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	userCount := math.MaxInt64
	userCount = 1000
	ch := make(chan bool, 10)
	for i := 0; i < userCount; i++ {
		wg.Add(1)
		ch <- true
		go func(ch chan bool, i int) {
			fmt.Println("go func:", i)
			defer wg.Done()
			time.Sleep(time.Second)
			<-ch
		}(ch, i)
	}
	wg.Wait()
}
