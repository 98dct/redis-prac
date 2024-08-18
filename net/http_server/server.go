package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", ":8888")
	if err != nil {
		panic(err)
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("接连接失败", err)
			continue
		}
		fmt.Printf("连接成功,来自%v\n", conn.RemoteAddr().String())

		go readWrite(conn)
	}
}

func readWrite(conn net.Conn) {
	defer conn.Close()

	reader := bufio.NewReader(conn)
	for {
		var buf [1024]byte
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("服务端读取失败！ ", err)
			return
		}
		got := string(buf[:n])
		fmt.Println("接收的数据:", got)
	}
}
