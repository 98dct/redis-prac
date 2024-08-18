package main

import (
	"fmt"
	"net"
	"time"
)

// 监听模式
func test1() {
	localAddr := &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 2000,
	}
	conn, err := net.ListenUDP("udp", localAddr)
	if err != nil {
		return
	}

	defer conn.Close()

	_, err = conn.WriteToUDP([]byte("world"), &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 2888,
	})
	if err != nil {
		panic(err)
	}

}

func test2() {
	localAddr := &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 3000,
	}
	conn, err := net.ListenUDP("udp", localAddr)
	if err != nil {
		return
	}

	defer conn.Close()

	_, err = conn.WriteToUDP([]byte("world"), &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 2000,
	})
	if err != nil {
		panic(err)
	}

	time.Sleep(1000)
	data := make([]byte, 1024)
	n, _ := conn.Read(data)
	fmt.Printf("收到的消息:%s\n", data[:n])
}

func main() {
	//test1()
	test2()
}
