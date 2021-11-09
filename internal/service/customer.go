package service

import (
	"github.com/google/uuid"
	"github.com/sadetskayfu/rest-golang/internal/model"
)

// CREATE
func (s *Service) CreateCustomer(u *model.Customer) (*model.Customer, error) {

	res, err := s.repo.CreateCustomer(u)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GET ALL
func (s *Service) GetAllCustomer() ([]*model.Customer, error) {
	res, err := s.repo.GetAllCustomer()
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GET BY ID
func (s *Service) GetByIdCustomer(id uuid.UUID) (*model.Customer, error) {
	res, err := s.repo.GetByIdCustomer(id)
	if err != nil {
		panic(err)
	}
	return res, nil
}

// DELETE BY ID
func (s *Service) DeleteByIdCustomer(id uuid.UUID) (string, error) {
	res, err := s.repo.DeleteByIdCustomer(id)
	if err != nil {
		panic(err)
	}
	return res, nil
}

// UPDATE BY ID
func (s *Service) UpdateByIdCustomer(id uuid.UUID, u *model.Customer) (*model.Customer, error) {
	res, err := s.repo.UpdateByIdCustomer(id, u)
	if err != nil {
		panic(err)
	}
	return res, nil
}
