package main

import (
	"fmt"
	"os"
	"sync"
	"math/rand"
	"runtime/trace"
)

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	err = trace.Start(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer trace.Stop()
	var total int
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 100; j++ {
				total += readNumber()
			}
			wg.Done()
		}()
	}

	wg.Wait()
}

//go:noinline
func readNumber() int {
	return rand.Intn(10)
}
