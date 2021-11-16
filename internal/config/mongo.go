// Package config for connect to DB
package config

// Return string for connect to mongo DB

func GetMongoConnectionString() string {
	connectString := "mongodb+srv://Rest_Api_Golang:12345@cluster0.fslpu.mongodb.net/Customers?retryWrites=true&w=majority"
	return connectString
}
