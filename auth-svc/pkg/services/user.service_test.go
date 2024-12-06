package services_test

import (
	"context"
	"errors"
	"go-microservices-grpc/auth-svc/pkg/data/models"
	"go-microservices-grpc/auth-svc/pkg/data/repositories"
	"go-microservices-grpc/auth-svc/pkg/pb"
	"go-microservices-grpc/auth-svc/pkg/services"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/types/known/emptypb"
)

func TestSignUpUser(t *testing.T) {
	mockRepo := new(repositories.MockUserRepository)

	mockRepo.On("Create", mock.AnythingOfType("*models.UserModel")).Return(nil)

	service := &services.UserService{
		Repository: mockRepo,
	}


	in := &pb.SignUpRequest{
		Email: "test@gmail.com",
		Password: "password",
	}
	res, err := service.SignUp(context.Background(), in)

	assert.Nil(t, err)
	assert.NotNil(t, res)
	mockRepo.AssertExpectations(t)
}

func TestSignInUserShouldSuccess(t *testing.T) {
	mockRepo := new(repositories.MockUserRepository)
	userStub := models.NewUserModel("name", "test@gmail.com", "P@$$w0rd")
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("P@$$w0rd"), bcrypt.DefaultCost)
	userStub.Password = string(hashedPassword)
	mockRepo.On("GetByEmail", mock.AnythingOfType("string")).Return(userStub, nil)

	service := &services.UserService{
		Repository: mockRepo,
	}


	in := &pb.SignInRequest{
		Email: "test@gmail.com",
		Password: "P@$$w0rd",
	}
	res, err := service.SignIn(context.Background(), in)

	assert.Nil(t, err)
	assert.NotNil(t, res)
	assert.Equal(t, userStub.Email, res.Email)
}

func TestSignInUserShouldFail(t *testing.T) {
	mockRepo := new(repositories.MockUserRepository)
	userStub := models.NewUserModel("name", "test@gmail.com", "P@$$w0rd")
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("P@$$w0rd"), bcrypt.DefaultCost)
	userStub.Password = string(hashedPassword)
	mockRepo.On("GetByEmail", mock.AnythingOfType("string")).Return(userStub, nil)

	service := &services.UserService{
		Repository: mockRepo,
	}


	in := &pb.SignInRequest{
		Email: "test@gmail.com",
		Password: "wrongpassword",
	}
	_, err := service.SignIn(context.Background(), in)

	assert.NotNil(t, err)
	assert.EqualError(t, err, "Invalid email or password")
}

func TestUpdateUser(t *testing.T){
	mockRepo := new(repositories.MockUserRepository)

	userStub := models.NewUserModel("name", "test@gmail.com", "password")
	mockRepo.On("GetById", mock.AnythingOfType("string")).Return(userStub, nil)
	mockRepo.On("UpdateUser", mock.AnythingOfType("*models.UserModel")).Return(userStub, nil)

	service := &services.UserService{
		Repository: mockRepo,
	}

	input := &pb.User{
		Id: userStub.ID.Hex(),
		Name: "changed",
		Email: "test@gmail.com",
	}

	user, err := service.UpdateUser(context.Background(), input)

	assert.Nil(t, err)
	assert.Equal(t, user.Name, "changed")
	mockRepo.AssertExpectations(t)
}

func TestValidateTokenError(t *testing.T){
	mockRepo := new(repositories.MockUserRepository)

	
	mockRepo.On("GetById", mock.AnythingOfType("string")).Return(nil, errors.New("error"))

	service := &services.UserService{
		Repository: mockRepo,
	}

	input := &pb.ValidateTokenRequest{
	Token: "token",
	}

	user, err := service.ValidateToken(context.Background(), input)

	assert.Nil(t, user)
	assert.NotNil(t, err)
	mockRepo.AssertExpectations(t)
}

func TestValidateToken(t *testing.T){
	mockRepo := new(repositories.MockUserRepository)

	userStub := models.NewUserModel("name", "test@gmail.com", "password")
	mockRepo.On("GetById", mock.Anything).Return(userStub, nil)
	
	service := &services.UserService{
		Repository: mockRepo,
	}

	input := &pb.ValidateTokenRequest{
		Token: userStub.ID.Hex(),
	}

	user, err := service.ValidateToken(context.Background(), input)
	
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, userStub.Email, user.Email)
	mockRepo.AssertExpectations(t)
}




func TestDeleteUser(t *testing.T){
	mockRepo := new(repositories.MockUserRepository)
	
	mockRepo.On("DeleteUser", mock.AnythingOfType("string")).Return(nil)

	service := &services.UserService{
		Repository: mockRepo,
	}

	input := &pb.DeleteUserRequest{
		Id: primitive.NewObjectID().Hex(),
	}

	user, err := service.DeleteUser(context.Background(), input)

	assert.Nil(t, err)
	assert.Equal(t, user, &emptypb.Empty{})
	mockRepo.AssertExpectations(t)
}


