package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", ":8888")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	fmt.Printf("连接服务端成功！:%v\n", conn.RemoteAddr())

	_, err = conn.Write([]byte("Hello"))
	if err != nil {
		panic(err)
	}

	//time.Sleep(1 * time.Second)
	// 直接关闭，不经过四次挥手
	conn.(*net.TCPConn).SetLinger(0)
	conn.Close()

	fmt.Println("客户端已关闭,")
	time.Sleep(60 * time.Second)
}
