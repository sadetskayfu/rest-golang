package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/sadetskayfu/rest-golang/internal/model"
	"go.mongodb.org/mongo-driver/bson"

)

type CatMongoDB struct{
	DataBaseName string
	CollectionName string
}
var newCatMongoDB = CatMongoDB{
	DataBaseName: "rest",
	CollectionName: "cats",

}
// GET ALL 
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

// GET BY ID
func (r *MongoRepository) GetByIdCat(id uuid.UUID)(*model.Cat,error){
	filter := bson.D{{"_id",id}}
	u:= &model.Cat{}
	res := r.DB.Database(newCatMongoDB.DataBaseName).Collection(newCatMongoDB.CollectionName).FindOne(context.TODO(),filter)
	err := res.Decode(&u)
	if err !=nil {
		return nil,err
	}
	return u,nil
}

// UPDATE BY ID
func (r *MongoRepository) UpdateByIdCat(id uuid.UUID, u *model.Cat)(*model.Cat,error){
	filter := bson.D{{"_id",id}}
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
	u.Id = id
	return u,nil
}

// DELETE BY ID
func (r *MongoRepository) DeleteByIdCat(id uuid.UUID)(string,error){
	filter := bson.D{{"_id",id}}
	_,err := r.DB.Database(newCatMongoDB.DataBaseName).Collection(newCatMongoDB.CollectionName).DeleteOne(context.TODO(),filter)
	if err != nil {
		panic(err)
	}
	res := "Cat deleted"
	return res,nil
}

// CREATE
func (r *MongoRepository) CreateCat(u *model.Cat)(*model.Cat,error){
	u.Id = uuid.New()
	_,err := r.DB.Database(newCatMongoDB.DataBaseName).Collection(newCatMongoDB.CollectionName).InsertOne(context.TODO(),u)
	if err != nil {
		return nil,err
	}
	return u,nil
}