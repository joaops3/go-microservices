package repositories

import (
	"go-microservices-grpc/auth-svc/pkg/data/models"

	"github.com/stretchr/testify/mock"
)

type MockUserRepository struct {
	mock.Mock
}


func (m *MockUserRepository) Create(data *models.UserModel) error {
	args := m.Called(data)
	return args.Error(0)
}


func (m *MockUserRepository) GetByEmail(email string) (*models.UserModel, error) {
	args := m.Called(email)
	return args.Get(0).(*models.UserModel),args.Error(0)

}

func (m *MockUserRepository) GetById(id string) (*models.UserModel, error) {

	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	userStub := models.NewUserModel("name", "test@gmail.com", "password")
	return userStub, nil
}

func (m *MockUserRepository) UpdateUser(data *models.UserModel) (*models.UserModel, error) {
	args := m.Called(data)
	return args.Get(0).(*models.UserModel), nil
}


func (m *MockUserRepository) DeleteUser(id string) error {
	args := m.Called(id)
	return args.Error(0)
}