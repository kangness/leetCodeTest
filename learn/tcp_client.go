package main

import (
	"fmt"
	"net"
)
//TcpClient.go

func main() {
	var msg = []byte("abcd")
	conn, err := net.Dial("tcp", "127.0.0.1:8080")
	if err != nil {
		panic(err)
	}
	// 发送消息
	_, err = conn.Write([]byte(msg))
	if err != nil {
		panic(err)
	}
	fmt.Println("发送消息: " + string(msg))
	// 读取消息
	_, err = conn.Read(msg[:])
	if err != nil {
		panic(err)
	}
	fmt.Println("收到消息: " + string(msg))
}