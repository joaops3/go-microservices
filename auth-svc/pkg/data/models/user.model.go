package models

import (
	"go-microservices-grpc/auth-svc/pkg/data/db"

	"time"

	"go-microservices-grpc/auth-svc/pkg/pb"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

const COLLECTION_USER string = "users"

type UserModel struct {
	*BaseModel `bson:",inline"`
	Name      string `json:"name" bson:"name,omitempty"  `
	Email      string `json:"email" bson:"email,omitempty"  `
	Password   string `json:"password" bson:"password,omitempty"`
	
}

type JwtResponse struct {
	Id   string `json:"_id"`
	Token string `json:"token"`
}

func GetDbUserCollection() *mongo.Collection {
	Db := db.GetDb()
	model := Db.Collection(COLLECTION_USER)
	return model
}

func NewUserModel(name string, email string, password string) *UserModel {
	v := &UserModel{Name: name, Email: email, Password: password, BaseModel: &BaseModel{}}
	v.InitBaseModel()
	return v
}


func (u *UserModel) ToProtoBuffer() *pb.User {
	return &pb.User{
		Id:       u.ID.Hex(),
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
		CreatedAt:  u.CreatedAt.Unix(),
		UpdatedAt:  u.UpdatedAt.Unix(),
	}
}
	
func (u *UserModel) FromProtoBuffer(user *pb.User) {
	u.ID, _ = primitive.ObjectIDFromHex(user.GetId())
	u.Name = user.GetName()
	u.Email = user.GetEmail()
	t := time.Unix(user.UpdatedAt, 0)
	u.UpdatedAt = &t
	u.CreatedAt = time.Unix(user.CreatedAt, 0)
	
}