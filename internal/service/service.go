// package service to execute logical processes
package service

import (
	"github.com/sadetskayfu/rest-golang/internal/repository"
	"github.com/sadetskayfu/rest-golang/pkg/logging"
)

// Service ...

type Service struct {
	repo repository.Repository
	log *logging.Logger
}

// NewService ...

func NewService(repo repository.Repository, log *logging.Logger) *Service {
	return &Service{
		repo: repo,
		log : log,
	}
}
