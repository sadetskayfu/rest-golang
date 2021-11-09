package handler

import (
	"github.com/sadetskayfu/rest-golang/internal/service"
)

type Handler struct {
	srv *service.Service
}

func NewHandler(srv *service.Service) *Handler {
	return &Handler{
		srv: srv,
	}
}


