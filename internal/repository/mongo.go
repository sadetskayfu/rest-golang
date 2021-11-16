package repository

import (
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"github.com/sadetskayfu/rest-golang/logging/logger"
)

// MongoRepository ...

type MongoRepository struct {
	DB *mongo.Client	
}

// NewMongoRepository ...

func NewMongoRepository(client *mongo.Client) *MongoRepository{
	logger.Logger.log.info("Connect to database MONGO")
	return &MongoRepository{
		DB : client,
	}
}
