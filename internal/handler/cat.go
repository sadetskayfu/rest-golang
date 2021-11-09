package handler

import (
	"net/http"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sadetskayfu/rest-golang/internal/model"
)

// CREATE
func (h *Handler) CreateCat(c echo.Context) error {
	u := &model.Cat{}
	err := c.Bind(u)
	if err != nil {
		panic(err)
	}
	res, err := h.srv.CreateCat(u)
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusCreated, res)
}

// GET ALL
func (h *Handler) GetAllCat(c echo.Context) error {
	res, err := h.srv.GetAllCat()
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK, res)
}

// GET BY ID
func (h *Handler) GetByIdCat(c echo.Context) error {
	id := uuid.Must(uuid.Parse(c.Param("id")))
	res, err := h.srv.GetByIdCat(id)
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK, res)
}

// DELETE BY ID
func (h *Handler) DeleteByIdCat(c echo.Context) error {
	id := uuid.Must(uuid.Parse(c.Param("id")))
	res, err := h.srv.DeleteByIdCat(id)
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK, res)
}

// UPDATE BY ID
func (h *Handler) UpdateByIdCat(c echo.Context) error {
	u := &model.Cat{}
	err := c.Bind(u)
	if err != nil {
		panic(err)
	}
	id := uuid.Must(uuid.Parse(c.Param("id")))
	res, err := h.srv.UpdateByIdCat(id, u)
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusCreated, res)
}
