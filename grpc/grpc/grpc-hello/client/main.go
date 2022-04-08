package main

import (
	"context"
	"grpc-demo/proto"
	"log"
	"os"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	ADDRESS      = "localhost:50051"
	DEFAULT_NAME = "world"
)

func main() {
	// 建立连接
	conn, err := grpc.Dial(ADDRESS, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalln("failed to connect: ", err)
	}
	defer conn.Close()
	c := proto.NewGreeterClient(conn)

	// contact the server and print out its response
	name := DEFAULT_NAME
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	// 1s 的上下文
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &proto.HelloRequest{Name: name})
	if err != nil {
		log.Fatalln("failed to greet: ", err)
	}
	log.Printf("Greeting: %s", r.Message)
}
