package model

import "github.com/google/uuid"

type Customer struct {
	Id      uuid.UUID `json:"Id" bson:"_id"`
	Name    string    `json:"Name" bson:"Name"`
	Email   string    `json:"Email" bson:"Email"`
	Age     int       `json:"Age" bson:"Age"`
	Married bool      `json:"Married" bson:"Married"`
}
