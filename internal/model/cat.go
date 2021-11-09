package model

import (
	"github.com/google/uuid"
)

type Cat struct {
	Id    uuid.UUID `json:"Id" bson:"_id"`
	Name  string    `json:"Name" bson:"Name"`
	Age   int       `json:"Age" bson:"Age"`
	Color string    `json:"Color" bson:"Color"`
}

