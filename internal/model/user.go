package model

import (
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

type User struct {
	Id       uuid.UUID `json:"Id" bson:"_id"`
	Email    string    `json:"Email" bson:"Email"`
	Password string    `json:"Password" bson:"Password"`
	Name     string    `json:"Name" bson:"Name"`
	Age      int       `json:"Age" bson:"Age"`
}

type AuthUser struct {
	Id       uuid.UUID `json:"Id" bson:"_id"`
	Email    string    `json:"Email" bson:"Email"`
	Password string    `json:"Password" bson:"Password"`
}

type JwtCustomClaims struct {
	Id    uuid.UUID `json:"Id" bson:"_id"`
	jwt.StandardClaims
}
