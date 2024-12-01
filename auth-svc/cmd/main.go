package main

import (
	"context"
	"go-microservices-grpc/auth-svc/pkg/data/db"
	"go-microservices-grpc/auth-svc/pkg/pb"
	"go-microservices-grpc/auth-svc/pkg/services"
	"os"

	"net"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	err := godotenv.Load()
	
	if err != nil {
	  panic(err.Error())
	}
	
	dbClient, err := db.InitDb()
	if err != nil {
		panic(err)
	}
 
	defer dbClient.Disconnect(context.Background())

	

	userService := services.NewUserService()

	grpcServer := grpc.NewServer()
	pb.RegisterAuthServiceServer(grpcServer, userService)

	reflection.Register(grpcServer)

	PORT := os.Getenv("PORT")
	listen, err := net.Listen("tcp", PORT)

	if err != nil {
		panic(err)
	}

	if err := grpcServer.Serve(listen); err != nil {
		panic(err)
	}
}