// package service to execute logical processes
package service

import (
	"github.com/sadetskayfu/rest-golang/internal/repository"
)

// Service ...

type Service struct {
	repo repository.Repository
}

// NewService ...

func NewService(repo repository.Repository) *Service {
	return &Service{
		repo: repo,
	}
}
