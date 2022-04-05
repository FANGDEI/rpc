package main

import (
	"context"
	"grpc-demo/proto"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	PORT = ":50051"
)

// 服务对象
type server struct {
	proto.UnimplementedGreeterServer
}

// SayHello 实现服务的接口 在 proto 中定义的所有服务都是接口
func (s *server) SayHello(ctx context.Context, request *proto.HelloRequest) (*proto.HelloResponse, error) {
	return &proto.HelloResponse{Message: "hello " + request.Name}, nil
}

func main() {
	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalln("failed to listen: ", err)
	}

	// 起一个服务
	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &server{})
	// 注册反射服务 这个服务是CLI使用的 跟服务本身没有关系
	reflection.Register(s)
	if err := s.Serve(listener); err != nil {
		log.Fatalln("failed to serve: ", err)
	}
}
