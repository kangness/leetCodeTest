package main 
import (
    //"runtime"
    "time"
)

func hello() {
}
func main( ) {
     //runtime.GOMAXPROCS(1)
     go func() {
           for {
			hello()
		}
     }()
     time.Sleep(time.Millisecond)
     println("OK")
}
