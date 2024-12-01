package db

import (
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
)

var	( 
	client *mongo.Client 
	
)



func InitDb() (*mongo.Client, error){
	var err error
	
	client, err = InitializeMongo()

	if err != nil {
		return nil, fmt.Errorf("error initializing mongo: %v", err)
	}
	
	return client, nil
}


func GetDb() *mongo.Database {
	nameDb := os.Getenv("MONGO_DB_NAME")


	db := client.Database(nameDb) 

	if db == nil {
		panic("Error getting db")
	}

	return db
}