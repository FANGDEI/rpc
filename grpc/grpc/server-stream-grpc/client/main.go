package main

import (
	"context"
	"log"
	"server-stream-grpc/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	ADDRESS = "localhost"
	PORT = ":50051"
)

func main() {
	conn, err := grpc.Dial(ADDRESS + PORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("failed to connection: ", err)
	}
	defer conn.Close()
	
	c := proto.NewGreeterClient(conn)
	request := &proto.StreamRequest{
		Data: "aaa",
	}
	response, _ := c.GetStream(context.Background(), request)
	for {
		aa, err := response.Recv()
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(aa)
	}
}