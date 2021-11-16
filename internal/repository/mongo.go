package repository

import (
	"github.com/sadetskayfu/pkg/logging/logging"
	"go.mongodb.org/mongo-driver/mongo"
)

// MongoRepository ...

type MongoRepository struct {
	DB *mongo.Client	
}

// NewMongoRepository ...

func NewMongoRepository(client *mongo.Client) *MongoRepository{
	logging.Logger.info("HGFSDFSDFSDF")
	
	return &MongoRepository{
		DB : client,
	}
}
