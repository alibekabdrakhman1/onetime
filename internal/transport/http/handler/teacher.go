package handler

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
	"onTime/internal/model"
	"onTime/internal/service"
	"onTime/internal/transport/http/middleware"
)

type TeacherHandler struct {
	Service *service.Service
	jwt     *middleware.JWTAuth
}

func (h *TeacherHandler) Login(c echo.Context) error {
	var input model.LogIn
	if err := c.Bind(&input); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	fmt.Println(input)

	teacher, err := h.Service.Teacher.Login(c.Request().Context(), input)
	if err != nil {
		fmt.Println(err)
		return err
	}
	token, err := h.jwt.GenerateJWT(teacher.Login)
	c.Set("Teacher", teacher)
	return c.JSON(http.StatusOK, token)
}

func (h *TeacherHandler) SignUp(c echo.Context) error {
	var input model.TeacherCreate

	if err := c.Bind(&input); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	fmt.Println(input)
	teacher := model.Teacher{
		Login:    input.Login,
		Password: input.Password,
		Name:     input.Name,
		Classes:  []string{},
	}
	id, err := h.Service.Teacher.SignUp(c.Request().Context(), teacher)

	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *TeacherHandler) GetByLogin(c echo.Context) error {
	teacher, err := h.Service.Teacher.GetByLogin(c.Request().Context(), c.Param("login"))
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, teacher)
}

func (h *TeacherHandler) GetClasses(c echo.Context) error {
	login := c.Request().Context().Value(model.ContextData{}).(string)
	classes, err := h.Service.Teacher.GetClasses(c.Request().Context(), login)
	if err != nil {
		fmt.Println(err)
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, classes)
}

func (h *TeacherHandler) CreateClass(c echo.Context) error {
	login := c.Request().Context().Value(model.ContextData{}).(string)
	var input model.CreateClass
	if err := c.Bind(&input); err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}
	createClass, err := h.Service.Teacher.CreateClass(c.Request().Context(), input, login)
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, createClass)
}

type ITeacherHandler interface {
	SignUp(c echo.Context) error
	GetByLogin(c echo.Context) error
	GetClasses(c echo.Context) error
	CreateClass(c echo.Context) error
	Login(c echo.Context) error
}

func NewTeacherHandler(s *service.Service, jwt *middleware.JWTAuth) *TeacherHandler {
	return &TeacherHandler{
		Service: s,
		jwt:     jwt,
	}
}
