package service

import (
	"github.com/google/uuid"
	"github.com/sadetskayfu/rest-golang/internal/model"
)

// CreateCustomer service ...

func (s *Service) CreateCustomer(u *model.Customer) (*model.Customer, error) {
	res, err := s.repo.CreateCustomer(u)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetAllCustomer service...

func (s *Service) GetAllCustomer() ([]*model.Customer, error) {
	res, err := s.repo.GetAllCustomer()
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetByIdCustomer service...

func (s *Service) GetByIDCustomer(ID uuid.UUID) (*model.Customer, error) {
	res, err := s.repo.GetByIDCustomer(ID)
	if err != nil {
		panic(err)
	}
	return res, nil
}

// DeleteByIdCustomer service...

func (s *Service) DeleteByIDCustomer(ID uuid.UUID) (string, error) {
	res, err := s.repo.DeleteByIDCustomer(ID)
	if err != nil {
		panic(err)
	}
	return res, nil
}

// UpdateByIdCustomer service...

func (s *Service) UpdateByIDCustomer(ID uuid.UUID, u *model.Customer) (*model.Customer, error) {
	res, err := s.repo.UpdateByIDCustomer(ID, u)
	if err != nil {
		panic(err)
	}
	return res, nil
}
