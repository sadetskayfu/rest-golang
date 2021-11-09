package service

import (
	"github.com/google/uuid"
	"github.com/sadetskayfu/rest-golang/internal/model"
)

// CREATE
func (s *Service) CreateCat(u *model.Cat) (*model.Cat, error) {
	res, err := s.repo.CreateCat(u)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GET ALL
func (s *Service) GetAllCat() ([]*model.Cat, error) {
	res, err := s.repo.GetAllCat()
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GET BY ID
func (s *Service) GetByIdCat(id uuid.UUID) (*model.Cat, error) {
	res, err := s.repo.GetByIdCat(id)
	if err != nil {
		panic(err)
	}
	return res, nil
}

// DELETE BY ID
func (s *Service) DeleteByIdCat(id uuid.UUID) (string, error) {
	res, err := s.repo.DeleteByIdCat(id)
	if err != nil {
		panic(err)
	}
	return res, nil
}

// UPDATE BY ID
func (s *Service) UpdateByIdCat(id uuid.UUID, u *model.Cat) (*model.Cat, error) {
	res, err := s.repo.UpdateByIdCat(id, u)
	if err != nil {
		panic(err)
	}
	return res, nil
}
