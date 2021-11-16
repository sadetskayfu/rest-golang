package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/sadetskayfu/rest-golang/internal/model"
	"go.mongodb.org/mongo-driver/bson"
)

// Struct variables for use table in mongo DB

type UserMongoDB struct {
	DataBaseName   string
	CollectionName string
}

// Variable for use table in mongo DB

var newUserMongoDB = CatMongoDB{
	DataBaseName:   "rest",
	CollectionName: "users",
}

// Registration user in mongo DB...

func (r *MongoRepository) Registration(u *model.User) (*model.User, error) {
	u.ID = uuid.New()
	_, err := r.DB.Database(newUserMongoDB.DataBaseName).Collection(newUserMongoDB.CollectionName).InsertOne(context.TODO(), u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// LogIn user in mongo DB...

func (r *MongoRepository) LogIn(u *model.AuthUser) (*model.AuthUser, error) {
	filter := bson.D{
		{"Email", u.Email},
		{"Password", u.Password}}
	res := r.DB.Database(newUserMongoDB.DataBaseName).Collection(newUserMongoDB.CollectionName).FindOne(context.TODO(), filter)
    err := res.Decode(&u)
	if err != nil {
		return nil,err
	}
	return u, nil
}
