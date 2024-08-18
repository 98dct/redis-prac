package main

import (
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"net"
	"redis-prac/net/proto/user"
	"time"
)

type UserService struct {
	user.UnimplementedUserServiceServer
}

func (s *UserService) GetUserInfo(ctx context.Context, r *user.UserRequest) (*user.UserResponse, error) {
	fmt.Printf("try to get user info by %s", r.GetName())
	time.Sleep(2 * time.Second)
	if len(r.GetName()) == 0 {
		return nil, errors.New("empty query name!")
	}

	return &user.UserResponse{
		Id:       123,
		Username: r.GetName(),
		Nickname: "尊敬的" + r.GetName(),
	}, nil
}

func main() {

	l, err := net.Listen("tcp", "127.0.0.1:8090")
	if err != nil {
		fmt.Println("建立监听出错")
		return
	}

	// 实例化grpc客户端
	server := grpc.NewServer()
	// 注册UserService服务
	user.RegisterUserServiceServer(server, &UserService{})
	// 向grpc服务端注册反射服务
	reflection.Register(server)
	// 启动grpc服务
	if err := server.Serve(l); err != nil {
		fmt.Println("启动grpc服务失败！")
	}
}
