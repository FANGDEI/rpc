package main

import (
	"context"
	"fmt"
	"log"
	"server-client-stream-grpc/proto"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const ADDRESS = "localhost:50051"

func CheckError(err error) {
	if err != nil {
		log.Fatalln("something wrong")
	}
}

func main() {
	conn, err := grpc.Dial(ADDRESS, grpc.WithTransportCredentials(insecure.NewCredentials()))
	CheckError(err)
	defer conn.Close()

	client := proto.NewGreeterClient(conn)
	request, err := client.AllStream(context.Background())
	CheckError(err)
	for i := 0; i < 10; i++ {
		request.Send(&proto.StreamRequest{Data: fmt.Sprintf("this is %dth message!!!", i+1)})
		response, err := request.Recv()
		CheckError(err)
		fmt.Println("The message which is sended from server is received", response)
		time.Sleep(time.Second)
	}
}
