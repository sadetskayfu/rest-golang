package handler

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/sadetskayfu/rest-golang/internal/model"
)

// Registration handler...

func (h *Handler)Registration(c echo.Context) error {
	u := &model.User{}
	err := c.Bind(u)
	if err != nil {
		panic(err)
	}
	res, err := h.srv.Registration(u)
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusCreated, res)
}

// LogIn handler...

func (h *Handler) LogIn(c echo.Context) error {
	u := &model.AuthUser{}
	err := c.Bind(u)
	if err != nil {
		panic(err)
	}
	res, err := h.srv.LogIn(u)
	if err != nil {
		panic(err)
	}
	return c.JSON(http.StatusOK, echo.Map{
		"message": "You were logged in!",
		"token": res})
}
