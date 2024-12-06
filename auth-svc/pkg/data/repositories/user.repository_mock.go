package repositories

import (
	"go-microservices-grpc/auth-svc/pkg/data/models"

	"github.com/stretchr/testify/mock"
	"golang.org/x/crypto/bcrypt"
)

type MockUserRepository struct {
	mock.Mock
}


func (m *MockUserRepository) Create(data *models.UserModel) error {
	args := m.Called(data)
	return args.Error(0)
}


func (m *MockUserRepository) GetByEmail(email string) (*models.UserModel, error) {
	userStub := models.NewUserModel("name", "test@gmail.com", "P@$$w0rd")
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte("P@$$w0rd"), bcrypt.DefaultCost)
	userStub.Password = string(hashedPassword)
	return userStub, nil

}

func (m *MockUserRepository) GetById(id string) (*models.UserModel, error) {
	args := m.Called(id)
	
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	
	return args.Get(0).(*models.UserModel), args.Error(1)
}

func (m *MockUserRepository) UpdateUser(data *models.UserModel) (*models.UserModel, error) {
	args := m.Called(data)
	return args.Get(0).(*models.UserModel), nil
}


func (m *MockUserRepository) DeleteUser(id string) error {
	args := m.Called(id)
	return args.Error(0)
}