package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/sadetskayfu/rest-golang/internal/model"
	"go.mongodb.org/mongo-driver/bson"
)

type CustomerMongoDB struct {
	DataBaseName   string
	CollectionName string
}

var newCustomerMongoDB = CatMongoDB{
	DataBaseName:   "rest",
	CollectionName: "Customers",
}

// GET ALL
func (r *MongoRepository) GetAllCustomer() ([]*model.Customer, error) {
	filter := bson.M{}
	res, err := r.DB.Database(newCustomerMongoDB.DataBaseName).Collection(newCustomerMongoDB.CollectionName).Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	allCustomers := []*model.Customer{}
	for res.Next(context.TODO()) {
		u := &model.Customer{}
		err := res.Decode(&u)
		if err != nil {
			panic(err)
		}
		allCustomers = append(allCustomers, u)
	}
	res.Close(context.TODO())
	return allCustomers, nil

}

// GET BY ID
func (r *MongoRepository) GetByIdCustomer(id uuid.UUID) (*model.Customer, error) {
	filter := bson.D{{"_id", id}}
	u := &model.Customer{}
	res := r.DB.Database(newCustomerMongoDB.DataBaseName).Collection(newCustomerMongoDB.CollectionName).FindOne(context.TODO(), filter)
	err := res.Decode(&u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// UPDATE BY ID
func (r *MongoRepository) UpdateByIdCustomer(id uuid.UUID, u *model.Customer) (*model.Customer, error) {
	filter := bson.D{{"_id", id}}
	update := bson.D{
		{"$set", bson.D{
			{"Name", u.Name},
			{"Email", u.Email},
			{"Age", u.Age},
			{"Married", u.Married},
		}}}
	_, err := r.DB.Database(newCustomerMongoDB.DataBaseName).Collection(newCustomerMongoDB.CollectionName).UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}
	u.Id = id
	return u, nil
}

// DELETE BY ID
func (r *MongoRepository) DeleteByIdCustomer(id uuid.UUID) (string, error) {
	filter := bson.D{{"_id", id}}
	_, err := r.DB.Database(newCustomerMongoDB.DataBaseName).Collection(newCustomerMongoDB.CollectionName).DeleteOne(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	res := "Customer deleted"
	return res, nil
}

// CREATE
func (r *MongoRepository) CreateCustomer(u *model.Customer) (*model.Customer, error) {
	u.Id = uuid.New()
	_, err := r.DB.Database(newCustomerMongoDB.DataBaseName).Collection(newCustomerMongoDB.CollectionName).InsertOne(context.TODO(), u)
	if err != nil {
		return nil, err
	}
	return u, nil
}
