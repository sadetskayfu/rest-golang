package repository

import (
	"github.com/google/uuid"
	"github.com/sadetskayfu/rest-golang/internal/model"
)

// Customer interface ...

type Customer interface {
	CreateCustomer(*model.Customer) (*model.Customer, error)
	GetAllCustomer() ([]*model.Customer, error)
	GetByIDCustomer(uuid.UUID) (*model.Customer, error)
	DeleteByIDCustomer(uuid.UUID) (string, error)
	UpdateByIDCustomer(uuid.UUID, *model.Customer) (*model.Customer, error)
}

// Cat interface ...

type Cat interface {
	GetAllCat() ([]*model.Cat, error)
	GetByIDCat(uuid.UUID) (*model.Cat, error)
	DeleteByIDCat(uuid.UUID) (string, error)
	UpdateByIDCat(uuid.UUID, *model.Cat) (*model.Cat, error)
	CreateCat(*model.Cat) (*model.Cat, error)
}

// Authentication interface...

type Authentication interface {
	Registration(*model.User)(*model.User,error)
	LogIn(*model.AuthUser)(*model.AuthUser,error)
}


// Repository interface ...

type Repository interface {
	Customer
	Cat
	Authentication
}
