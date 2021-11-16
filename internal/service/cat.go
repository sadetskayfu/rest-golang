package service

import (
	"github.com/google/uuid"
	"github.com/sadetskayfu/rest-golang/internal/model"
)

// CreateCat service...

func (s *Service) CreateCat(u *model.Cat) (*model.Cat, error) {
	res, err := s.repo.CreateCat(u)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetAllCat service...

func (s *Service) GetAllCat() ([]*model.Cat, error) {
	res, err := s.repo.GetAllCat()
	if err != nil {
		return nil, err
	}
	return res, nil
}

// GetByIdCat service...

func (s *Service) GetByIDCat(ID uuid.UUID) (*model.Cat, error) {
	res, err := s.repo.GetByIDCat(ID)
	if err != nil {
		panic(err)
	}
	return res, nil
}

// DeleteByIdCat service...

func (s *Service) DeleteByIDCat(ID uuid.UUID) (string, error) {
	res, err := s.repo.DeleteByIDCat(ID)
	if err != nil {
		panic(err)
	}
	return res, nil
}

// UpdateByIdCat service...

func (s *Service) UpdateByIDCat(ID uuid.UUID, u *model.Cat) (*model.Cat, error) {
	res, err := s.repo.UpdateByIDCat(ID, u)
	if err != nil {
		panic(err)
	}
	return res, nil
}
