package handler

import (
	"net/http"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sadetskayfu/rest-golang/internal/model"
)

// CREATE
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

// GET ALL
func (h *Handler) GetAllCustomer(c echo.Context) error {
	res, err := h.srv.GetAllCustomer()
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK, res)
}

// GET BY ID
func (h *Handler) GetByIdCustomer(c echo.Context) error {

	id := uuid.Must(uuid.Parse(c.Param("id")))
	res, err := h.srv.GetByIdCustomer(id)
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK, res)
}

// DELETE BY ID
func (h *Handler) DeleteByIdCustomer(c echo.Context) error {
	id := uuid.Must(uuid.Parse(c.Param("id")))
	res, err := h.srv.DeleteByIdCustomer(id)
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK, res)
}

// UPDATE BY ID
func (h *Handler) UpdateByIdCustomer(c echo.Context) error {
	u := &model.Customer{}
	err := c.Bind(u)
	if err != nil {
		panic(err)
	}
	id := uuid.Must(uuid.Parse(c.Param("id")))
	res, err := h.srv.UpdateByIdCustomer(id, u)
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusCreated, res)
}
