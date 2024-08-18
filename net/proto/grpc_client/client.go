package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"redis-prac/net/proto/user"
	"time"
)

func main() {

	// 远程连接凭证，insecure模式下禁用了传输安全认证
	credentials := grpc.WithTransportCredentials(insecure.NewCredentials())
	// 连接grpc服务器
	conn, err := grpc.Dial("127.0.0.1:8090", credentials)
	if err != nil {
		fmt.Println("连接失败!", err)
		return
	}
	defer conn.Close()

	userCilent := user.NewUserServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, err := userCilent.GetUserInfo(ctx, &user.UserRequest{Name: "golang"})
	if err != nil {
		fmt.Println("调用getuserInfo方法失败！", err)
		return
	}

	fmt.Printf("用户Id: %+v\n", res.Id)
	fmt.Printf("用户name: %+v\n", res.Username)
	fmt.Printf("用户nickname: %+v\n", res.Nickname)
}
