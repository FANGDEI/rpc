package main

import (
	"fmt"
	"log"
	"net"
	"server-client-stream-grpc/proto"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const PORT = ":50051"

type Server struct {

}

func (s *Server) AllStream(streamRequest proto.Greeter_AllStreamServer) error {
	for i := 0; i < 10; i++ {
		request, err := streamRequest.Recv()
		CheckError(err)
		fmt.Println("The message which is sended from client is received: ", request)
		streamRequest.Send(&proto.StreamResponse{Data: "The server has receive the message: " + request.String()})
		time.Sleep(time.Second)
	}
	return nil
}

func CheckError(err error) {
	if err != nil {
		log.Fatalln("something wrong: ", err)
	}
}

func main() {
	listener, err := net.Listen("tcp", PORT)
	CheckError(err)

	s := grpc.NewServer()
	proto.RegisterGreeterServer(s, &Server{})
	reflection.Register(s)
	
	s.Serve(listener)
}