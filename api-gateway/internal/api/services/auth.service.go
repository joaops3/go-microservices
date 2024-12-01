package services

import (
	"context"

	"go-microservices-grpc/api-gateway/internal/pb"
)

type AuthService struct {
	pb.UnimplementedAuthServiceServer
}

type AuthServiceInterface interface {
	pb.AuthServiceServer
	SignUp(context.Context, *pb.SignUpRequest) (*pb.SignUpResponse, error)
    SignIn(context.Context, *pb.SignInRequest) (*pb.SignInResponse, error)

}

func NewAuthService() AuthServiceInterface {

	return &AuthService{}
}


func (r *AuthService)SignUp(context.Context, *pb.SignUpRequest) (*pb.SignUpResponse, error) {


	return nil, nil
}

func (r *AuthService) SignIn(context.Context, *pb.SignInRequest) (*pb.SignInResponse, error) {


	return nil, nil
}

