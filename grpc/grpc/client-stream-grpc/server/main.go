package main

import (
	"client-stream-grpc/proto"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const PORT = ":50051"

type Server struct {
}

func (s *Server) PutStream(putResponse proto.Greeter_PutStreamServer) error {
	for i := 0; i < 10; i++ {
		if item, err := putResponse.Recv(); err != nil {
			log.Fatalln("break, err: ", err)
		} else {
			log.Println(item)
		}
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
