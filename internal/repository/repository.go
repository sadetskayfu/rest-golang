package repository

import (
	"github.com/google/uuid"
	"github.com/sadetskayfu/rest-golang/internal/model"
)

// CUSTOMER INTERFACE
type Customer interface {
	CreateCustomer(*model.Customer) (*model.Customer, error)
	GetAllCustomer() ([]*model.Customer, error)
	GetByIdCustomer(uuid.UUID) (*model.Customer, error)
	DeleteByIdCustomer(uuid.UUID) (string, error)
	UpdateByIdCustomer(uuid.UUID, *model.Customer) (*model.Customer, error)
}

// CAT INTERFACE
type Cat interface {
	GetAllCat() ([]*model.Cat, error)
	GetByIdCat(uuid.UUID) (*model.Cat, error)
	DeleteByIdCat(uuid.UUID) (string, error)
	UpdateByIdCat(uuid.UUID, *model.Cat) (*model.Cat, error)
	CreateCat(*model.Cat) (*model.Cat, error)
}

type Authentication interface {
	Registration(*model.User)(*model.User,error)
	LogIn(*model.AuthUser)(*model.AuthUser,error)
}


// REPOSITORY INTERFACE
type Repository interface {
	Customer
	Cat
	Authentication
}
