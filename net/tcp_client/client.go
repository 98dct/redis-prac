package main

import (
	"fmt"
	"net"
	"time"
)

func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:1888")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	fmt.Printf("连接服务端成功！:%v\n", conn.RemoteAddr())

	_, err = conn.Write([]byte("Hello"))
	if err != nil {
		panic(err)
	}
	time.Sleep(10 * time.Millisecond)

	var buf [1021]byte
	n, err := conn.Read(buf[:])
	if err != nil {
		panic(err)
	}
	fmt.Println("收到客户端回复,", string(buf[:n]))
}
