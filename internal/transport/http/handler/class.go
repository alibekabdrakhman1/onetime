package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"onTime/internal/service"
)

type ClassHandler struct {
	Service *service.Service
}

func (h *ClassHandler) GetAllClasses(c echo.Context) error {
	classes, err := h.Service.Class.GetAllClasses(c.Request().Context())
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, classes)
}

func (h *ClassHandler) GetById(c echo.Context) error {
	class, err := h.Service.Class.GetById(c.Request().Context(), c.Param("id"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, class)
}

type IClassHandler interface {
	GetAllClasses(c echo.Context) error
	GetById(c echo.Context) error
}

func NewClassHandler(s *service.Service) *ClassHandler {
	return &ClassHandler{
		Service: s,
	}
}
