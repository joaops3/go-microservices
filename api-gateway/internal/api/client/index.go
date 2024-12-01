package client

import (
	"fmt"
	"go-microservices-grpc/api-gateway/internal/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ServiceClient struct {
	Client pb.AuthServiceClient
}

type Config struct {
	Port          string `mapstructure:"AUTH_PORT"`
	AuthSuvUrl    string `mapstructure:"AUTH_SVC_URL"`
}

func InitServiceClient(c *Config) pb.AuthServiceClient {
	fmt.Println("API Gateway :  InitServiceClient")
	//	using WithInsecure() because no SSL running
	cc, err := grpc.NewClient(c.AuthSuvUrl, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		fmt.Println("Could not connect:", err)
	}
	return pb.NewAuthServiceClient(cc)
}