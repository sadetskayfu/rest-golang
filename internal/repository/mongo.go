package repository

import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoRepository struct {
	DB *mongo.Client	
}

func NewMongoRepository(client *mongo.Client) *MongoRepository{
	fmt.Println("Connect to DATABASE mongo")
	return &MongoRepository{
		DB : client,
	}
}
