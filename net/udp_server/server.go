package main

import (
	"fmt"
	"net"
)

// 监听模式
// 此种模式下conn没有记录消息的来源地址
func test1() {
	localAddr := &net.UDPAddr{
		IP:   net.ParseIP("127.0.0.1"),
		Port: 2888,
	}
	conn, err := net.ListenUDP("udp", localAddr)
	if err != nil {
		return
	}

	defer conn.Close()

	data := make([]byte, 1024)
	for {
		n, remoteAddr, err := conn.ReadFromUDP(data)
		if err != nil {
			panic(err)
		}
		fmt.Printf("消息来源：%s, 消息内容：%s\n", remoteAddr, data[:n])
		_, err = conn.WriteToUDP([]byte("world"), &net.UDPAddr{
			IP:   net.ParseIP("127.0.0.1"),
			Port: 2000,
		})
		if err != nil {
			panic(err)
		}
	}
}

// 拨号模式
func test2() {
	destIp := net.ParseIP("127.0.0.1")
	localAddr := &net.UDPAddr{IP: net.IPv4zero, Port: 2000}
	destAddr := &net.UDPAddr{IP: destIp, Port: 3000}

	conn, err := net.DialUDP("udp", localAddr, destAddr)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	data := make([]byte, 1024)
	for {
		n, err := conn.Read(data)
		if err != nil {
			panic(err)
		}
		fmt.Printf("接收到的消息: %v\n", string(data[:n]))

		n, err = conn.Write([]byte("ok"))
		if err != nil {
			panic(err)
		}
	}

}

func main() {
	//test1()
	test2()
}
