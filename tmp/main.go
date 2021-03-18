package main

import (
	"context"
	"go-blog/proto"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err.Error())
	}
	client := proto.NewAuthServiceClient(conn)
	client.SignUp(context.Background(), &proto.SignUpRequest{Username: "Carl", Email: "carl@gmail.com", Password: "examplestring"})
}
