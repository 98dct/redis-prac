package main

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/rpc"
	"net/rpc/jsonrpc"
	"redis-prac/net/model"
)

type UserHandler struct {
}

func (u *UserHandler) GetUserInfo(name string, reply *model.UserDetail) error {
	fmt.Printf("try to get user info by %s", name)
	if len(name) == 0 {
		return errors.New("the query name cannot be empty")
	}

	reply.Id = 123
	reply.Name = name
	reply.NickName = "尊敬的" + name
	return nil
}

// http实现rpc通信
func test1() {
	rpc.Register(new(UserHandler))
	rpc.HandleHTTP()

	l, err := net.Listen("tcp", "127.0.0.1:8090")
	if err != nil {
		fmt.Println(err)
		return
	}

	err = http.Serve(l, nil)
	if err != nil {
		fmt.Println(err)
	}
}

// tcp实现rpc通信
func test2() {

	rpc.Register(new(UserHandler))
	l, err := net.Listen("tcp", "127.0.0.1:8090")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("建立tcp连接失败")
			continue
		}
		go func(conn net.Conn) {
			rpc.ServeConn(conn)
		}(conn)
	}

}

// jsonrpc 实现rpc通信
func test3() {
	rpc.Register(new(UserHandler))
	l, err := net.Listen("tcp", "127.0.0.1:8090")
	if err != nil {
		fmt.Println(err)
		return
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("建立tcp连接失败")
			continue
		}
		go func(conn net.Conn) {
			jsonrpc.ServeConn(conn)
		}(conn)
	}
}
func main() {
	//test1()
	//test2()
	test3()
}
