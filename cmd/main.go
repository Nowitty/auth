package main

import (
	userRepository "auth/internal/repository/user"
	userService "auth/internal/service/user"
	userApi "auth/internal/api/user"
	desc "auth/pkg/user_v1"
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"github.com/jackc/pgx/v4/pgxpool"
)

const grpcPort = 50051

const (
	dbDSN = "host=localhost port=54321 dbname=auth user=auth-user password=auth-password sslmode=disable"
)

func main() {
	ctx := context.Background()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", grpcPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	pool, err := pgxpool.Connect(ctx, dbDSN)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	defer pool.Close()

	userRepo := userRepository.NewRepository(pool)
	userSrv := userService.NewService(userRepo)

	s := grpc.NewServer()
	reflection.Register(s)
	desc.RegisterUserV1Server(s, userApi.NewImplementation(userSrv))

	log.Printf("Listening")

	if err = s.Serve(lis); err != nil {
		log.Fatalf("falied to serve")
	}
}
