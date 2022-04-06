package main

import (
	"fmt"
	"log"
	"net"
	"server-stream-grpc/proto"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const PORT = ":50051"

type Server struct {
}

func (s *Server) GetStream(request *proto.StreamRequest, response proto.Greeter_GetStreamServer) error {
	for i := 0; i < 10; i++ {
		response.Send(&proto.StreamResponse{
			Data: fmt.Sprintf("%v", time.Now().Unix()),
		})
		time.Sleep(1 * time.Second)
	}
	return nil
}

func main() {
	listener, err := net.Listen("tcp", PORT)
	if err != nil {
		log.Fatalln("failed to listen: ", err)
	}

	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &Server{})
	reflection.Register(s)
	s.Serve(listener)
}
