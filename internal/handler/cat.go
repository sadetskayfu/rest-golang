package handler

import (
	"net/http"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"github.com/sadetskayfu/rest-golang/internal/model"
)

// CreateCat handler...

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

//  GetAllCat handler...

func (h *Handler) GetAllCat(c echo.Context) error {
	res, err := h.srv.GetAllCat()
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK, res)
}

// GetByIdCat handler...

func (h *Handler) GetByIDCat(c echo.Context) error {
	ID := uuid.Must(uuid.Parse(c.Param("id")))
	res, err := h.srv.GetByIDCat(ID)
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK, res)
}

// DeleteByIdCat handler...

func (h *Handler) DeleteByIDCat(c echo.Context) error {
	ID := uuid.Must(uuid.Parse(c.Param("id")))
	res, err := h.srv.DeleteByIDCat(ID)
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK, res)
}

// UpdateByIdCat handler...

func (h *Handler) UpdateByIDCat(c echo.Context) error {
	u := &model.Cat{}
	err := c.Bind(u)
	if err != nil {
		panic(err)
	}
	ID := uuid.Must(uuid.Parse(c.Param("id")))
	res, err := h.srv.UpdateByIDCat(ID, u)
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusCreated, res)
}
