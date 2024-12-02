package services

import (
	"context"
	"errors"
	"go-microservices-grpc/auth-svc/pkg/data/models"
	"go-microservices-grpc/auth-svc/pkg/data/repositories"
	"go-microservices-grpc/auth-svc/pkg/pb"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/types/known/emptypb"
)

type UserService struct {
	Repository repositories.UserRepositoryInterface
	pb.UnimplementedAuthServiceServer
}

type UserServiceInterface interface {
	pb.AuthServiceServer
	SignUp(context.Context, *pb.SignUpRequest) (*pb.SignUpResponse, error)
    SignIn(context.Context, *pb.SignInRequest) (*pb.SignInResponse, error)
    UpdateUser(context.Context, *pb.User) (*pb.User, error)
    DeleteUser(context.Context, *pb.DeleteUserRequest) (*emptypb.Empty, error)
	ValidateToken(context.Context, *pb.ValidateTokenRequest) (*pb.User, error)
}

func NewUserService() UserServiceInterface {
	collection := models.GetDbUserCollection()

	return &UserService{
		Repository: repositories.NewUserRepository(collection),
	}
}

func (s *UserService) SignUp(ctx context.Context, in *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	user := models.NewUserModel(in.Name, in.Email, in.Password)
	hashed, err := bcrypt.GenerateFromPassword([]byte(in.Password), 10)
	if err != nil {
		return nil, err
	}
	user.Password = string(hashed)
	err = s.Repository.Create(user)

	if err != nil {
		return nil, err
	}
	
	resp := &pb.SignUpResponse{Name: user.Name, Email: user.Email }
	return resp, nil
}

func (s *UserService) SignIn(ctx context.Context, in *pb.SignInRequest) (*pb.SignInResponse, error) {
	found, err := s.Repository.GetByEmail(in.Email) 

	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(found.Password), []byte(in.Password))

	if err != nil {
		return nil, errors.New("Invalid email or password")
	}

	token, err := generateTokenJWT(found.ID.Hex())
	if err != nil {
		return nil, err
	}

	resp := models.JwtResponse{
		Token: token,
		Id: found.ID.Hex(),
	}

	return &pb.SignInResponse{Name: resp.Token, Email: in.Email }, nil

}


func (s *UserService) UpdateUser(ctx context.Context, in *pb.User) (*pb.User, error) {
	user, err := s.Repository.GetById(in.Id)
	if err != nil {
		return nil, err
	}

	if in.Name != "" {
		user.Name = in.Name
	}
	if in.Email != "" {
		user.Email = in.Email
	}
	if in.Password != "" {
		hashed, err := bcrypt.GenerateFromPassword([]byte(in.Password), 10)
		if err != nil {
			return nil, err
		}
		user.Password = string(hashed)
	}

	_, err = s.Repository.UpdateUser(user)
	if err != nil {
		return nil, err
	}

	return &pb.User{Id: user.ID.Hex(), Name: user.Name, Email: user.Email}, nil
}

func (s *UserService) ValidateToken(ctx context.Context, in *pb.ValidateTokenRequest) (*pb.User, error) {
	u, err := s.Repository.GetById(in.Token)
	if err != nil {
		return nil, err
	}
	if u == nil {
        return nil, errors.New("user not found")
    }

	return u.ToProtoBuffer(), nil
}

func (s *UserService) DeleteUser(ctx context.Context, in *pb.DeleteUserRequest) (*emptypb.Empty, error) {
	err := s.Repository.DeleteUser(in.Id)
	if err != nil {
		return nil, err
	}

	return &emptypb.Empty{}, nil
}


func generateTokenJWT(id string) (string, error){
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": id,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	
	
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		return "", err
	}

	return tokenString, err
}