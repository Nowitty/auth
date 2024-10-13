package main

import (
	auth "auth/pkg/user_v1"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const grpcPort = 50051

type server struct {
	auth.UnimplementedUserV1Server
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)
	auth.RegisterUserV1Server(s, &server{})

	log.Printf("Listening")

	if err = s.Serve(lis); err != nil {
		log.Fatalf("falied to serve")
	}
}
