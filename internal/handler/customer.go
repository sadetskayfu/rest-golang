package handler

import (
	"net/http"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sadetskayfu/rest-golang/internal/model"
)

// CreateCustomer handler...
func (h *Handler) CreateCustomer(c echo.Context) error {
	u := &model.Customer{}
	err := c.Bind(u)
	if err != nil {
		panic(err)
	}
	res, err := h.srv.CreateCustomer(u)
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusCreated, res)
}

// GetAllCustomer handler...
func (h *Handler) GetAllCustomer(c echo.Context) error {
	res, err := h.srv.GetAllCustomer()
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK, res)
}

// GetByIdCustomer handler...
func (h *Handler) GetByIdCustomer(c echo.Context) error {
	ID := uuid.Must(uuid.Parse(c.Param("id")))
	res, err := h.srv.GetByIDCustomer(ID)
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK, res)
}

// DeleteByIdCustomer handler...
func (h *Handler) DeleteByIDCustomer(c echo.Context) error {
	ID := uuid.Must(uuid.Parse(c.Param("id")))
	res, err := h.srv.DeleteByIDCustomer(ID)
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK, res)
}

// UpdateByIdCustomer handler...

func (h *Handler) UpdateByIDCustomer(c echo.Context) error {
	u := &model.Customer{}
	err := c.Bind(u)
	if err != nil {
		panic(err)
	}
	ID := uuid.Must(uuid.Parse(c.Param("id")))
	res, err := h.srv.UpdateByIDCustomer(ID, u)
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusCreated, res)
}
