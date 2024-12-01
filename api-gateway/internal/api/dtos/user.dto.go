package dtos

import (
	"go-microservices-grpc/api-gateway/internal/pb"

	validation "github.com/go-ozzo/ozzo-validation"
)

type CreateUserDto struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password"`
}

func (d *CreateUserDto) Validate() error {
	return validation.ValidateStruct(d,
		validation.Field(&d.Email, validation.Required.Error("O Email é obrigatório"), validation.Length(1, 50)),
		validation.Field(&d.Password, validation.Required.Error("A Senha é obrigatório"), validation.Length(1, 50)),
	)
}

func (s *CreateUserDto) ToProtoBuff() *pb.SignUpRequest {
	return &pb.SignUpRequest{
		Email:    s.Email,
		Password: s.Password,
	}
}

