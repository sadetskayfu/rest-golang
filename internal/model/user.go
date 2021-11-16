package model

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)
// User struct

type User struct {
	ID       uuid.UUID `json:"Id" bson:"_id"`
	Email    string    `json:"Email" bson:"Email"`
	Password string    `json:"Password" bson:"Password"`
	Name     string    `json:"Name" bson:"Name"`
	Age      int       `json:"Age" bson:"Age"`
}

// AuthUser struct

type AuthUser struct {
	ID      uuid.UUID `json:"Id" bson:"_id"`
	Email    string    `json:"Email" bson:"Email"`
	Password string    `json:"Password" bson:"Password"`
}

// JwtCustomClaims struct

type JwtCustomClaims struct {
	ID    uuid.UUID `json:"Id" bson:"_id"`
	jwt.StandardClaims
}
