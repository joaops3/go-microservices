package db

import (
	"context"
	"fmt"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitializeMongo() (*mongo.Client, error) {
	
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	MONGO_URL := os.Getenv("MONGO_URL")
	
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(MONGO_URL))

	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Printf("ERROR MONGO: %v", err.Error())
		client.Disconnect(ctx)
		panic(err)
	}
	
	if err != nil {
		fmt.Printf("ERROR MONGO: %v", err.Error())
		client.Disconnect(ctx)
		panic(err)
	}
	
	
	return client, nil
}