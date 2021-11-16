package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/sadetskayfu/rest-golang/internal/model"
	"go.mongodb.org/mongo-driver/bson"
)

// Struct variables for use table in mongo DB

type CatMongoDB struct{
	DataBaseName string
	CollectionName string
}

// Variable for use table in mongo DB

var newCatMongoDB = CatMongoDB{
	DataBaseName: "rest",
	CollectionName: "cats",
}

// GetAllCat in mongo DB... 

func (r *MongoRepository) GetAllCat()([]*model.Cat,error){
	filter := bson.M{}
	res,err := r.DB.Database(newCatMongoDB.DataBaseName).Collection(newCatMongoDB.CollectionName).Find(context.TODO(),filter)
	if err != nil {
		return nil,err
	}
	allCats := []*model.Cat{}
	for res.Next(context.TODO()){
		u:=&model.Cat{}
		err:=res.Decode(&u)
		if err != nil {
			panic (err)
		}
		allCats = append(allCats, u)
	}
	res.Close(context.TODO())
	return allCats,nil
	
}

// GetByIdCat in mongo DB...

func (r *MongoRepository) GetByIDCat(ID uuid.UUID)(*model.Cat,error){
	filter := bson.D{{"_id",ID}}
	u:= &model.Cat{}
	res := r.DB.Database(newCatMongoDB.DataBaseName).Collection(newCatMongoDB.CollectionName).FindOne(context.TODO(),filter)
	err := res.Decode(&u)
	if err !=nil {
		return nil,err
	}
	return u,nil
}

// UpdateByIdCat in mongo DB...

func (r *MongoRepository) UpdateByIDCat(ID uuid.UUID, u *model.Cat)(*model.Cat,error){
	filter := bson.D{{"_id",ID}}
	update := bson.D{
		{"$set", bson.D{
			{"Name", u.Name},
			{"Age", u.Age},
			{"Color", u.Color},
			}},}
			_, err := r.DB.Database(newCatMongoDB.DataBaseName).Collection(newCatMongoDB.CollectionName).UpdateOne(context.TODO(), filter, update)
			if err != nil {
				return nil, err
			}
	u.ID = ID
	return u,nil
}

// DeleteByIdCat in mongo DB...

func (r *MongoRepository) DeleteByIDCat(ID uuid.UUID)(string,error){
	filter := bson.D{{"_id",ID}}
	_,err := r.DB.Database(newCatMongoDB.DataBaseName).Collection(newCatMongoDB.CollectionName).DeleteOne(context.TODO(),filter)
	if err != nil {
		panic(err)
	}
	const res = "Cat deleted"
	return res,nil
}

// CreateCat in mongo DB...

func (r *MongoRepository) CreateCat(u *model.Cat)(*model.Cat,error){
	u.ID = uuid.New()
	_,err := r.DB.Database(newCatMongoDB.DataBaseName).Collection(newCatMongoDB.CollectionName).InsertOne(context.TODO(),u)
	if err != nil {
		return nil,err
	}
	return u,nil
}