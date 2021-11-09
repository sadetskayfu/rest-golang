package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/sadetskayfu/rest-golang/internal/model"
	"go.mongodb.org/mongo-driver/bson"
)

type UserMongoDB struct {
	DataBaseName   string
	CollectionName string
}

var newUserMongoDB = CatMongoDB{
	DataBaseName:   "rest",
	CollectionName: "users",
}

func (r *MongoRepository) Registration(u *model.User) (*model.User, error) {
	u.Id = uuid.New()
	_, err := r.DB.Database(newUserMongoDB.DataBaseName).Collection(newUserMongoDB.CollectionName).InsertOne(context.TODO(), u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

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
