package main
import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var lc = new(sync.Mutex)
	//这个locker 为啥传入一个引用？
	var cond = sync.NewCond(lc)
	for i := 0; i < 3; i++ {
		go func(x int) {
		   //竞争锁
			cond.L.Lock()
			//记得要释放锁
			defer cond.L.Unlock()
			cond.Wait()
			fmt.Println(x)
			cond.Signal()
		}(i)
	}
	//睡眠一会，确保下面的Signal()能通知到一个（难道可能通知不到？）
	fmt.Println("now start frist sleep")
	time.Sleep(2*time.Second)
	fmt.Println("now finish frist sleep")
	cond.Signal()
	fmt.Println("now finish signal")
	//cond.Broadcast()
	fmt.Println("now finish Broadcast")
	fmt.Println("now start second sleep")
	time.Sleep(2*time.Second)
	fmt.Println("now finish second sleep")
}
