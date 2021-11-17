package repository

import (
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

// MongoRepository ...

type MongoRepository struct {
	DB *mongo.Client	
}

// NewMongoRepository ...

func NewMongoRepository(client *mongo.Client) *MongoRepository{
	fmt.Println("Connected to db mongo")
	return &MongoRepository{
		DB : client,
	}
}
