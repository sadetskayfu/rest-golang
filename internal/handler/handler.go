package handler

import (
	"github.com/sadetskayfu/rest-golang/internal/service"
)

// Handler ...

type Handler struct {
	srv *service.Service
}

// NewHandler ...

func NewHandler(srv *service.Service) *Handler {
	return &Handler{
		srv: srv,
	}
}


