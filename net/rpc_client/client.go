package main

import (
	"fmt"
	"net/rpc"
	"net/rpc/jsonrpc"
	"redis-prac/net/model"
)

// http实现rpc通信
func test1() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:8090")
	if err != nil {
		fmt.Println(err)
		return
	}

	user := new(model.UserDetail)
	err = client.Call("UserHandler.GetUserInfo", "golang", user)
	if err != nil {
		fmt.Println("用户信息出错")
		return
	}
	fmt.Printf("用户id: %+v\n", user.Id)
	fmt.Printf("用户name: %+v\n", user.Name)
	fmt.Printf("用户昵称: %+v\n", user.NickName)
}

// tcp实现rpc通信
func test2() {

	client, err := rpc.Dial("tcp", "127.0.0.1:8090")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer client.Close()

	user := new(model.UserDetail)
	err = client.Call("UserHandler.GetUserInfo", "golang", user)
	if err != nil {
		fmt.Println("获得用户信息出错", err.Error())
		return
	}
	fmt.Printf("用户id: %+v\n", user.Id)
	fmt.Printf("用户name: %+v\n", user.Name)
	fmt.Printf("用户昵称: %+v\n", user.NickName)

}

// jsonrpc实现rpc通信
func test3() {
	client, err := jsonrpc.Dial("tcp", "127.0.0.1:8090")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer client.Close()

	user := new(model.UserDetail)
	err = client.Call("UserHandler.GetUserInfo", "golang", user)
	if err != nil {
		fmt.Println("获得用户信息出错", err.Error())
		return
	}
	fmt.Printf("用户id: %+v\n", user.Id)
	fmt.Printf("用户name: %+v\n", user.Name)
	fmt.Printf("用户昵称: %+v\n", user.NickName)
}
func main() {

	//test1()
	test2()

}
