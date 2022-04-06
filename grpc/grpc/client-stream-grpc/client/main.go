package main

import (
	"client-stream-grpc/proto"
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	ADDRESS = "localhost"
	PORT    = ":50051"
)

func CheckError(err error) {
	if err != nil {
		log.Fatalln("something wrong")
	}
}

func main() {
	conn, err := grpc.Dial(ADDRESS+PORT, grpc.WithTransportCredentials(insecure.NewCredentials()))
	CheckError(err)
	defer conn.Close()

	c := proto.NewGreeterClient(conn)
	putRequest, _ := c.PutStream(context.Background())
	for i := 0; i < 10; i++ {
		putRequest.Send(&proto.StreamRequest{
			Data: "ss",
		})
		time.Sleep(time.Second)
	}
}
