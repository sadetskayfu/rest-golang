// Package model for empaty struct

package model


import (
	"github.com/google/uuid"
)

// Cat struct
type Cat struct {
	ID    uuid.UUID `json:"Id" bson:"_id"`
	Name  string    `json:"Name" bson:"Name"`
	Age   int       `json:"Age" bson:"Age"`
	Color string    `json:"Color" bson:"Color"`
}

