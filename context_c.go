package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"
)

func handlerMessage(ctx context.Context, i int) {
	select {
	case <- ctx.Done():
		fmt.Println("this is ctx.done", time.Now(), ctx.Err())
		return
	default:
		st := int64(rand.Int() % 5000)
		fmt.Println("start to do this", i, st)
		time.Sleep(time.Microsecond * time.Duration(st))
		fmt.Println("end to do this", i)
	}
	fmt.Println(time.Now())
}
func main() {
	rand.Seed(time.Now().UnixNano())
	ctx , _ := context.WithTimeout(context.Background(), time.Second * 3)
	for i := 0; i < 10; i ++ {
		go handlerMessage(ctx, i)
	}
	fmt.Println("start to sleep ")
	time.Sleep(time.Second * 5)
	fmt.Println("end to sleep")
}
