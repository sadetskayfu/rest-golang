package model

import "github.com/google/uuid"

// Customer struct

type Customer struct {
	ID     uuid.UUID `json:"Id" bson:"_id"`
	Name    string    `json:"Name" bson:"Name"`
	Email   string    `json:"Email" bson:"Email"`
	Age     int       `json:"Age" bson:"Age"`
	Married bool      `json:"Married" bson:"Married"`
}
