package repository

import (
	"context"
	"github.com/google/uuid"
	"github.com/sadetskayfu/rest-golang/internal/model"
	"go.mongodb.org/mongo-driver/bson"
)

// Struct variables for use table in mongo DB

type CustomerMongoDB struct {
	DataBaseName   string
	CollectionName string
}

// Variable for use table in mongo DB

var newCustomerMongoDB = CatMongoDB{
	DataBaseName:   "rest",
	CollectionName: "Customers",
}

// GetAllCustomer in mongo db...

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

// GetByIdCustomer in mongo db...

func (r *MongoRepository) GetByIDCustomer(ID uuid.UUID) (*model.Customer, error) {
	filter := bson.D{{"_id", ID}}
	u := &model.Customer{}
	res := r.DB.Database(newCustomerMongoDB.DataBaseName).Collection(newCustomerMongoDB.CollectionName).FindOne(context.TODO(), filter)
	err := res.Decode(&u)
	if err != nil {
		return nil, err
	}
	return u, nil
}

// UpdateByIdCustomer in mongo db...

func (r *MongoRepository) UpdateByIDCustomer(ID uuid.UUID, u *model.Customer) (*model.Customer, error) {
	filter := bson.D{{"_id", ID}}
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
	u.ID = ID
	return u, nil
}

// DeleteByIdCustomer in mongo db...

func (r *MongoRepository) DeleteByIDCustomer(ID uuid.UUID) (string, error) {
	filter := bson.D{{"_id", ID}}
	_, err := r.DB.Database(newCustomerMongoDB.DataBaseName).Collection(newCustomerMongoDB.CollectionName).DeleteOne(context.TODO(), filter)
	if err != nil {
		panic(err)
	}
	const res = "Customer deleted"
	return res, nil
}

// CreateCustomer in mongo db...

func (r *MongoRepository) CreateCustomer(u *model.Customer) (*model.Customer, error) {
	u.ID = uuid.New()
	_, err := r.DB.Database(newCustomerMongoDB.DataBaseName).Collection(newCustomerMongoDB.CollectionName).InsertOne(context.TODO(), u)
	if err != nil {
		return nil, err
	}
	return u, nil
}
