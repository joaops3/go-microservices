package repositories

import (
	"context"
	"go-microservices-grpc/auth-svc/pkg/data/models"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository struct {
	ModelDb *mongo.Collection
	
}

type UserRepositoryInterface interface {
	Create(user *models.UserModel) error
	GetByEmail(email string) (*models.UserModel, error)
	UpdateUser(user *models.UserModel) (*models.UserModel, error)
	DeleteUser(id string) error
}

func NewUserRepository(modelDb *mongo.Collection) UserRepositoryInterface {
	return &UserRepository{ModelDb: modelDb}
}


func (r *UserRepository) Create(user *models.UserModel) error {
	_, err := r.ModelDb.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}
	return nil
}


func (r *UserRepository) GetByEmail(email string) (*models.UserModel, error) {

	data := &models.UserModel{}

	err := r.ModelDb.FindOne(context.Background(), bson.D{{Key: "email", Value: email }}).Decode(data)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}
		return nil, err
	}

	return data, nil
}

func (r *UserRepository) UpdateUser(user *models.UserModel) (*models.UserModel, error) {
	filter := bson.D{{Key: "_id", Value: user.ID}}
	update := bson.D{{Key: "$set", Value: user}}
	t := time.Now()
	user.UpdatedAt = &t
	_, err := r.ModelDb.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return nil, err
	}
	return user, nil
}


func (r *UserRepository) DeleteUser(id string) error {
	_, err := r.ModelDb.DeleteOne(context.Background(), bson.D{{Key: "_id", Value: id}})
	if err != nil {
		return err
	}
	return nil
}